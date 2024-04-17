package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const standartTemplate string = "templates/standard.txt"

func main() {
	flag.Parse()
	result, err := run(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, result)
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func run(args []string) (string, error) {
	lines, err := convertString(args...)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(&lines), nil
}

func convertString(args ...string) (Lines, error) {
	lines := Lines{}
	if len(args) == 0 {
		return lines, nil
	}
	text := strings.Join(args, "")
	if text == "" {
		return lines, nil
	}

	if !isASCII(text) {
		return lines, errors.New("only ASCII characters are allowed")
	}

	templateLines, err := openTemplateFile()
	if err != nil {
		return lines, err
	}
	if len(templateLines) > 0 {
		text = strings.ReplaceAll(text, "\\n", "\n")
		byteArray := []byte(text)
		line := Line{}
		var lastRune rune
		for _, v := range byteArray {
			lastRune = rune(v)
			if v == '\n' {
				lines.addLine(line)
				line = Line{}
				continue
			}
			letter := templateLines[v-32]
			line.addSymbol(letter)
		}
		if lastRune != '\n' {
			lines.addLine(line)
		}
	}

	return lines, nil
}

func openTemplateFile() ([]Symbol, error) {
	template, err := os.Open(standartTemplate)
	symbols := make([]Symbol, 0)
	if err != nil {
		return symbols, err
	}
	defer template.Close()

	sc := bufio.NewScanner(template)

	symbol := Symbol{}
	for sc.Scan() {
		if sc.Text() == "" {
			if symbol.len() != 0 {
				symbols = append(symbols, symbol)
				symbol = Symbol{}
			}
		} else {
			symbol.addSymbolPart(sc.Text())
		}
	}

	if err := sc.Err(); err != nil {
		return make([]Symbol, 0), err
	}
	return symbols, nil
}
