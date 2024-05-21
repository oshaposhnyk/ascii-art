package ascii

import (
	"embed"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Config struct {
	Text     string
	Template string
}

func NewConfig(text string) Config {
	return Config{
		Text:     text,
		Template: "standard",
	}
}

//go:embed standard.txt
//go:embed shadow.txt
//go:embed thinkertoy.txt
var f embed.FS

const Offset = 32

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func Run(c Config, w io.Writer) error {
	lines, err := ConvertString(c)
	if err != nil {
		return err
	}
	fmt.Fprint(w, fmt.Sprint(&lines))
	return nil
}

func ConvertString(c Config) (Lines, error) {
	lines := Lines{}
	if c.Text == "" {
		return lines, nil
	}
	if c.Text == "\n" {
		lines.addLine(Line{})
		return lines, nil
	}

	if !IsASCII(c.Text) {
		return lines, ErrNoASCII
	}

	tChars, err := OpenTemplateFile(c.Template, c.Text)
	if err != nil {
		return lines, err
	}
	if len(tChars) > 0 {
		text := strings.ReplaceAll(c.Text, "\\n", "\n")
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
			letter := tChars[v]
			line.addSymbol(letter)
		}
		if lastRune != '\n' {
			lines.addLine(line)
		}
	}

	return lines, nil
}

func OpenTemplateFile(tName, text string) (map[byte]Symbol, error) {
	result := map[byte]Symbol{}
	tPath := tName + ".txt"
	content, err := f.ReadFile(tPath)
	if err != nil {
		return result, ErrInvalideTemplate
	}
	tCont := strings.ReplaceAll(string(content), "\r\n", "\n")
	sText := strings.Split(tCont, "\n\n")
	for i, sym := range sText {
		if strings.ContainsRune(text, rune(i+Offset)) {
			result[byte(i+Offset)] = CreateFromStr(sym)
		}
	}
	return result, nil
}
