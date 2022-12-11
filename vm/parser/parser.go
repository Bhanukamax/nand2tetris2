package parser

import (
	"bufio"
	"strings"
)

type parser struct {
	text    string
	lenght  int
	current string
	lines   []string
	ptr     int
}

func NewParser(input string) *parser {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return &parser{input, len(lines), "", lines, 0}
}

func (p *parser) advance() {
	if p.hasMoreLines() {
		p.ptr = p.ptr + 1
		p.current = p.lines[p.ptr-1]
	}
}

func (p *parser) Current() string {
	return p.current
}

func (p *parser) hasMoreLines() bool {
	return p.ptr <= p.lenght
}

func (p *parser) Parse() {
	 p.advance()
}
