package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const minPerThread = 25

// ── Параллельная обработка ──────────────────────────────────────────────────

// workerFunc — тип функции-обработчика чанка строк.
type workerFunc func(chunk []string) []string

// parallelWork делит входной слайс на чанки и обрабатывает их параллельно,
// аналогично C++ parallel_work с min_per_thread = 25.
func parallelWork(data []string, fn workerFunc) []string {
	length := len(data)
	if length == 0 {
		return nil
	}

	maxThreads := (length + minPerThread - 1) / minPerThread
	hw := runtime.NumCPU()
	if hw == 0 {
		hw = 4
	}
	numThreads := hw
	if maxThreads < numThreads {
		numThreads = maxThreads
	}

	blockSize := length / numThreads

	results := make([][]string, numThreads)
	var wg sync.WaitGroup

	for i := 0; i < numThreads-1; i++ {
		wg.Add(1)
		chunk := data[i*blockSize : (i+1)*blockSize]
		idx := i
		go func() {
			defer wg.Done()
			results[idx] = fn(chunk)
		}()
	}
	// последний блок — остаток
	wg.Add(1)
	go func() {
		defer wg.Done()
		results[numThreads-1] = fn(data[(numThreads-1)*blockSize:])
	}()

	wg.Wait()

	// собираем результаты в том же порядке
	var out []string
	for _, r := range results {
		out = append(out, r...)
	}
	return out
}

// ── Обработчики чанков ──────────────────────────────────────────────────────

// parallelParseData фильтрует строки: оставляет только те, что выглядят
// как диапазон IP (аналог get_parsing_data в C++).
func parallelParseData(chunk []string) []string {
	var out []string
	for _, line := range chunk {
		// убираем символы не из множества "0123456789.-"
		cleaned := cleanLine(line)
		if isValidRangeLine(cleaned) {
			out = append(out, cleaned)
		}
	}
	return out
}

// parallelToCIDR конвертирует строки "ip_start-ip_end" в CIDR-нотацию.
func parallelToCIDR(chunk []string) []string {
	var out []string
	for _, line := range chunk {
		parts := strings.SplitN(line, "-", 2)
		if len(parts) != 2 {
			continue
		}
		startIP := net.ParseIP(strings.TrimSpace(parts[0]))
		endIP := net.ParseIP(strings.TrimSpace(parts[1]))
		if startIP == nil || endIP == nil {
			continue
		}
		start4 := startIP.To4()
		end4 := endIP.To4()
		if start4 == nil || end4 == nil {
			continue
		}
		startU := binary.BigEndian.Uint32(start4)
		endU := binary.BigEndian.Uint32(end4)
		cidrs := rangeToCIDR(startU, endU)
		out = append(out, cidrs...)
	}
	return out
}

// parallelToRange фильтрует строки "ip_start-ip_end", валидируя оба IP.
func parallelToRange(chunk []string) []string {
	var out []string
	for _, line := range chunk {
		parts := strings.SplitN(line, "-", 2)
		if len(parts) != 2 {
			continue
		}
		if net.ParseIP(strings.TrimSpace(parts[0])) != nil &&
			net.ParseIP(strings.TrimSpace(parts[1])) != nil {
			out = append(out, line)
		}
	}
	return out
}

// ── IP-утилиты ───────────────────────────────────────────────────────────────

// cleanLine удаляет из строки символы, не входящие в "0123456789.-".
func cleanLine(s string) string {
	var b strings.Builder
	for _, c := range s {
		if c >= '0' && c <= '9' || c == '.' || c == '-' {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// rangeToCIDR рекурсивно конвертирует диапазон [ipStart, ipEnd] в список CIDR.
// Аналог range_boundaries_to_cidr из C++.
func rangeToCIDR(ipStart, ipEnd uint32) []string {
	bits := uint32(1)
	mask := uint32(1)
	var newIP uint32

	for bits < 32 {
		newIP = ipStart | mask
		if newIP > ipEnd || (ipStart>>bits)<<bits != ipStart {
			bits--
			mask >>= 1
			break
		}
		bits++
		mask = (mask << 1) | 1
	}

	newIP = ipStart | mask
	prefix := 32 - bits

	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipStart)
	cidr := fmt.Sprintf("%s/%d", ip.String(), prefix)

	cidrs := []string{cidr}
	if newIP < ipEnd {
		cidrs = append(cidrs, rangeToCIDR(newIP+1, ipEnd)...)
	}
	return cidrs
}

// ── Файловый вывод ───────────────────────────────────────────────────────────

// saveToFile сохраняет строки в файл, по одной на строку.
func saveToFile(lines []string, filename string) error {
	if len(lines) == 0 {
		return fmt.Errorf("nothing to save: slice is empty")
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file %s: %w", filename, err)
	}
	defer f.Close()

	sb := strings.Builder{}
	for _, line := range lines {
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	if _, err := f.WriteString(sb.String()); err != nil {
		return fmt.Errorf("write file %s: %w", filename, err)
	}

	info, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("Size of %s: %s\n", filename, humanReadable(uint64(info.Size())))
	}
	return nil
}

// ── Основная логика ──────────────────────────────────────────────────────────

// runInit скачивает данные для кода и сохраняет два файла: _RANGE и _CIDR.
func runInit(code, outputPath string, codeMap map[string]string, tp TypeParsing) error {
	// нормализуем путь
	if !strings.HasSuffix(outputPath, "/") {
		outputPath += "/"
	}
	// убираем всё после первой точки (аналог C++ erase от pos)
	if idx := strings.Index(outputPath[1:], "."); idx != -1 {
		outputPath = outputPath[:idx+1]
	}

	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outputPath, err)
	}

	data, err := parsingSite(code, tp)
	if err != nil {
		return fmt.Errorf("parsing_site(%s): %w", code, err)
	}

	name := codeMap[code]
	base := filepath.Join(outputPath, name)

	fmt.Printf("Start for    --> %s\n", name)
	fmt.Printf("Path to save --> %s\n", base)

	// RANGE
	rangeData := parallelWork(data, parallelToRange)
	if err := saveToFile(rangeData, base+"_RANGE.txt"); err != nil {
		return err
	}

	// CIDR
	cidrData := parallelWork(data, parallelToCIDR)
	if err := saveToFile(cidrData, base+"_CIDR.txt"); err != nil {
		return err
	}

	fmt.Printf("Saved two files --> %s_RANGE.txt and %s_CIDR.txt\n", base, base)
	return nil
}

// ── Вспомогательное ──────────────────────────────────────────────────────────

