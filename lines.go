package main

import "fmt"

type Symbol struct {
	symbolPart []string
}

type Line struct {
	symbols []Symbol
}

type Lines struct {
	lines []Line
}

const symbolHeigh int = 8

func (s *Symbol) addSymbolPart(part string) {
	if len(s.symbolPart) < symbolHeigh {
		s.symbolPart = append(s.symbolPart, part)
	}
}

func (s *Symbol) getPart(index int) string {
	return s.symbolPart[index]
}

func (s *Symbol) len() int {
	return len(s.symbolPart)
}

func (Line) makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (l *Line) addSymbol(symbol Symbol) {
	l.symbols = append(l.symbols, symbol)
}

func (l *Line) String() string {
	formatted := ""
	sHeigh := l.makeRange(0, symbolHeigh)
	if len(l.symbols) != 0 {
		for i, v := range sHeigh {
			for _, s := range l.symbols {
				formatted += s.getPart(v)
			}
			if i != len(sHeigh)-1 {
				formatted += "\n"
			}
		}
	}
	return fmt.Sprintln(formatted)
}

func (ls *Lines) addLine(l Line) {
	ls.lines = append(ls.lines, l)
}

func (ls *Lines) String() string {
	formatted := ""
	for _, l := range ls.lines {
		formatted += fmt.Sprint(&l)
	}
	return formatted
}
