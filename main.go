package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const standartTemplate string = "templates/standard.txt"

func main() {
	flag.Parse()
	run()
}

func run() {
	lines, err := convertString(flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s", &lines)
}

func convertString(args ...string) (Lines, error) {
	lines := Lines{}
	if len(args) == 0 {
		return lines, nil
	}

	templateLines, err := openTemplateFile()
	if err != nil {
		return lines, err
	}
	if len(templateLines) > 0 {
		text := strings.Join(args, "")
		text = strings.ReplaceAll(text, "\\n", "\n")
		byteArray := []byte(text)
		line := Line{}
		for _, v := range byteArray {
			if v == '\n' {
				lines.addLine(line)
				line = Line{}
				continue
			}
			letter := templateLines[v-32]
			line.addSymbol(letter)
		}
		lines.addLine(line)
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
