package main

import (
	"bufio"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Config struct {
	text     string
	template string
}

func NewConfig(text string) Config {
	return Config{
		text:     text,
		template: "standard",
	}
}

//go:embed standard.txt
//go:embed shadow.txt
//go:embed thinkertoy.txt
var f embed.FS

func main() {
	flag.Parse()
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	input := args[0]
	config := NewConfig(input)

	if len(args) > 1 {
		config.template = args[1]
	}
	err := run(config, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func run(c Config, w io.Writer) error {
	lines, err := convertString(c)
	if err != nil {
		return err
	}
	fmt.Fprint(w, fmt.Sprint(&lines))
	return nil
}

func convertString(c Config) (Lines, error) {
	lines := Lines{}
	if c.text == "" {
		return lines, nil
	}

	if !isASCII(c.text) {
		return lines, errors.New("only ASCII characters are allowed")
	}

	tChars, err := openTemplateFile(c.template)
	if err != nil {
		return lines, err
	}
	if len(tChars) > 0 {
		text := strings.ReplaceAll(c.text, "\\n", "\n")
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
			letter := tChars[v-32]
			line.addSymbol(letter)
		}
		if lastRune != '\n' {
			lines.addLine(line)
		}
	}

	return lines, nil
}

func openTemplateFile(tName string) ([]Symbol, error) {
	symbols := make([]Symbol, 0)
	tPath := tName + ".txt"
	template, err := f.Open(tPath)
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
			continue
		}
		symbol.addSymbolPart(sc.Text())
	}

	if symbol.len() != 0 {
		symbols = append(symbols, symbol)
	}

	if err := sc.Err(); err != nil {
		return make([]Symbol, 0), err
	}
	return symbols, nil
}
