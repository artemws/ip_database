package main

import (
	"fmt"
	"os"
)

const helpOut = `Usage: ipsuip [OPTIONS]
To show this help: --help OR -h

The program downloads the IP addresses of the selected country or mainland.

OPTIONS:
  -m, --mainland [CODE]   Without argument: prints mainland-code list.
                          With CODE: download IPs for that mainland.
  -c, --country [CODE]    Without argument: prints country-code list.
                          With CODE: download IPs for that country.
  -o, --output PATH       Directory where files with IP addresses will be saved.
                          To avoid errors, the path should contain only Latin
                          characters without spaces or special shell characters.
      --all               Downloads entire database for all countries OR mainlands.
                          Can take a very long time!

EXAMPLES:
  1) ipsuip -c US -o /home/user/Desktop/
  2) ipsuip -c GE -o ./
  3) ipsuip --all -c -o ~/Downloads/CountryIPs/
  4) ipsuip -m EU -o /home/user/Desktop/Europe
`

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print(helpOut)
		os.Exit(0)
	}

	outputPath := "./"
	code := ""
	tp := TypeCountry
	all := false
	typeSet := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--help", "-h":
			fmt.Print(helpOut)
			os.Exit(0)

		case "-c", "--country":
			tp = TypeCountry
			typeSet = true
			// следующий аргумент — код, если он не начинается с '-'
			if i+1 < len(args) && args[i+1][0] != '-' {
				i++
				code = args[i]
			} else {
				// без аргумента — вывести список стран
				PrintMap(MapCode(CountryCodeData))
				os.Exit(0)
			}

		case "-m", "--mainland":
			tp = TypeContinent
			typeSet = true
			if i+1 < len(args) && args[i+1][0] != '-' {
				i++
				code = args[i]
			} else {
				PrintMap(MapCode(MainlandCodeData))
				os.Exit(0)
			}

		case "-o", "--output":
			if i+1 < len(args) {
				i++
				outputPath = args[i]
			} else {
				fmt.Fprintln(os.Stderr, "Error: -o requires a path argument")
				os.Exit(1)
			}

		case "--all":
			all = true

		default:
			fmt.Fprintf(os.Stderr, "Unknown argument: %s\n", arg)
			fmt.Print(helpOut)
			os.Exit(1)
		}
	}

	if !typeSet && !all {
		fmt.Fprintln(os.Stderr, "Error: specify -c (country) or -m (mainland)")
		fmt.Print(helpOut)
		os.Exit(1)
	}

	if err := run(all, tp, code, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(all bool, tp TypeParsing, code, outputPath string) error {
	if all {
		if tp == TypeCountry {
			m := MapCode(CountryCodeData)
			for c := range m {
				if err := runInit(c, outputPath, m, TypeCountry); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
				}
			}
		} else {
			m := MapCode(MainlandCodeData)
			for c := range m {
				if err := runInit(c, outputPath, m, TypeContinent); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
				}
			}
		}
		return nil
	}

	// одиночный запрос
	if code == "" {
		return fmt.Errorf("no country/mainland code specified")
	}

	if tp == TypeCountry {
		if _, ok := CodesCountries[code]; !ok {
			return fmt.Errorf("unknown country code %q", code)
		}
		return runInit(code, outputPath, MapCode(CountryCodeData), TypeCountry)
	}

	if _, ok := CodesMainlands[code]; !ok {
		return fmt.Errorf("unknown mainland code %q", code)
	}
	return runInit(code, outputPath, MapCode(MainlandCodeData), TypeContinent)
}
