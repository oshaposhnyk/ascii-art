package main

import "fmt"

type Symbol struct {
	s []string
}

type Line struct {
	s []Symbol
}

type Lines struct {
	l []Line
}

const symbolHeigh int = 8

func (s *Symbol) addSymbolPart(part string) {
	if len(s.s) < symbolHeigh {
		s.s = append(s.s, part)
	}
}

func (s *Symbol) getPart(index int) string {
	return s.s[index]
}

func (s *Symbol) len() int {
	return len(s.s)
}

func (Line) makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (l *Line) addSymbol(symbol Symbol) {
	l.s = append(l.s, symbol)
}

func (l *Line) String() string {
	formatted := ""
	sHeigh := l.makeRange(0, symbolHeigh)
	if len(l.s) != 0 {
		for i, v := range sHeigh {
			for _, s := range l.s {
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
	ls.l = append(ls.l, l)
}

func (ls *Lines) String() string {
	formatted := ""
	for _, l := range ls.l {
		formatted += fmt.Sprint(&l)
	}
	return formatted
}
