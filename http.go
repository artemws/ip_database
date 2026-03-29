package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// httpGet выполняет HTTPS GET запрос к host+target и возвращает строки тела ответа.
func httpGet(host, target string) ([]string, error) {
	url := "https://" + host + target

	client := &http.Client{
		Timeout: 5 * time.Minute, // большой таймаут — ответы могут быть объёмными
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", "ipsuip/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	var lines []string
	sc := bufio.NewScanner(resp.Body)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024) // буфер 1 МБ на строку
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan body: %w", err)
	}

	bytes := 0
	for _, l := range lines {
		bytes += len(l) + 1
	}
	fmt.Printf("Downloaded %s\n", humanReadable(uint64(bytes)))

	if len(lines) == 0 {
		return nil, fmt.Errorf("empty response from %s", url)
	}
	return lines, nil
}

// humanReadable форматирует байты в человекочитаемый вид (аналог C++ HumanReadable).
func humanReadable(size uint64) string {
	const units = "BKMGTPE"
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := uint64(1024), 0
	n := size / div
	for n >= 1024 && exp < len(units)-2 {
		div *= 1024
		exp++
		n = size / div
	}
	mantissa := float64(size) / float64(div)
	return fmt.Sprintf("%.1f%cB (%d)", mantissa, units[exp+1], size)
}

// parsingSite скачивает данные с suip.biz для кода страны или материка.
func parsingSite(code string, tp TypeParsing) ([]string, error) {
	var target string
	switch tp {
	case TypeContinent:
		target = "/?act=all-country-ip&continent=" + code + "&all-download"
	case TypeCountry:
		target = "/?act=all-country-ip&country=" + code + "&all-download"
	}

	lines, err := httpGet("suip.biz", target)
	if err != nil {
		return nil, err
	}

	result := parallelWork(lines, parallelParseData)
	if len(result) == 0 {
		return nil, fmt.Errorf("buffer empty after parsing for code %s", code)
	}
	return result, nil
}

// TypeParsing — тип географической выборки.
type TypeParsing int

const (
	TypeCountry   TypeParsing = iota
	TypeContinent TypeParsing = iota
)

// isValidRangeLine проверяет, что строка выглядит как "X.X.X.X-Y.Y.Y.Y"
// (начинается и заканчивается цифрой, содержит только цифры, точки и дефисы).
func isValidRangeLine(s string) bool {
	if len(s) < 15 {
		return false
	}
	if len(s) == 0 || s[0] < '0' || s[0] > '9' {
		return false
	}
	if s[len(s)-1] < '0' || s[len(s)-1] > '9' {
		return false
	}
	for _, c := range s {
		if !strings.ContainsRune("0123456789.-", c) {
			return false
		}
	}
	return true
}
