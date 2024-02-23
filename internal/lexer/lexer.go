package lexer

import (
	"fmt"
	"strings"

	"unicode/utf8"

	"github.com/pxlctzn/monkey/internal/token"
)

const EOF = -1

type lexer struct {
	input      string
	atEOF      bool
	startPos   int
	currentPos int
	linePos    int
	charPos    int
	lineSizes  []int
	ch         chan token.Token
}

func New(s string) chan token.Token {
	lxr := lexer{
		input:      s,
		ch:         make(chan token.Token),
		startPos:   0,
		currentPos: 0,
		linePos:    1,
		charPos:    0,
		lineSizes:  []int{0},
	}

	go lxr.run()

	return lxr.ch
}

func (l *lexer) errorf(format string, v ...any) stateFn {
	l.ch <- token.New(token.ILLEGAL, fmt.Sprintf(format, v...), l.linePos, l.startPos)

	return nil
}

func (l *lexer) acceptUntil(delimiter string) (c int) {
	for r := l.next(); !strings.ContainsRune(delimiter, r); r = l.next() {
		c += 1
	}

	l.backup()

	return c - 1
}

func (l *lexer) backup() {
	r, s := utf8.DecodeLastRuneInString(l.input[:l.currentPos])
	if isEndOfLine(r) {
		l.linePos -= 1
		l.charPos = l.lineSizes[l.linePos]
	} else {
		l.charPos -= 1
	}
	l.currentPos -= s

}
func (l *lexer) accept(valid string) (c int) {
	c = 0

	for r := l.next(); strings.ContainsRune(valid, r); r = l.next() {
		c++
	}
	l.backup()
	return c - 1
}

func (l *lexer) emit(typ token.Type) {
	l.ch <- token.New(typ, l.input[l.startPos:l.currentPos], l.linePos, l.startPos)

	l.startPos = l.currentPos
}

// Consume the next rune
func (l *lexer) next() rune {
	if l.currentPos >= len(l.input) {
		l.atEOF = true
		return EOF
	}

	r, s := utf8.DecodeRuneInString(l.input[l.currentPos:])
	if isEndOfLine(r) {
		l.linePos = l.linePos + 1
		l.lineSizes = append(l.lineSizes, l.charPos+1)
		l.charPos = 0
	} else {
		l.charPos = l.charPos + 1
	}
	l.currentPos = l.currentPos + s

	return r
}

func (l *lexer) peek() rune {
	if l.currentPos >= len(l.input) {
		l.atEOF = true
		return EOF
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.currentPos:])
	return r
}

func (l *lexer) run() {
	for state := parseNextState; state != nil; state = state(l) {
	}
	close(l.ch)
}
