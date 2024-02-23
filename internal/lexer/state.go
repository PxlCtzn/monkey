package lexer

import (
	"fmt"

	"github.com/pxlctzn/monkey/internal/token"
)

type stateFn func(*lexer) stateFn

const (
	space      = " "
	digits     = "0123456789"
	identifier = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	assign     = '='
	plus       = '+'
	minus      = '-'
	bang       = '!'
	asterisk   = '*'
	slash      = '/'
	lt         = '<'
	gt         = '>'
	comma      = ','
	semi_colon = ';'
	lparen     = '('
	rparen     = ')'
	lbrace     = '{'
	rbrace     = '}'
)

var symbols = map[rune]token.Type{
	assign:     token.ASSIGN,
	plus:       token.PLUS,
	minus:      token.MINUS,
	bang:       token.BANG,
	asterisk:   token.ASTERISK,
	slash:      token.SLASH,
	lt:         token.LT,
	gt:         token.GT,
	comma:      token.COMMA,
	semi_colon: token.SEMICOLON,
	lparen:     token.LPAREN,
	rparen:     token.RPAREN,
	lbrace:     token.LBRACE,
	rbrace:     token.RBRACE,
}
var keywords = map[string]token.Type{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"return": token.RETURN,
	"if":     token.IF,
	"else":   token.ELSE,
}

func parseNextState(l *lexer) stateFn {
	// fmt.Println(">> parseNextState")
	if l.atEOF {
		return nil
	}

	r := l.peek()

	if r == EOF {
		return nil
	}

	if r == '\n' {
		return newlineState
	}

	if isSpace(r) {
		return spaceState
	}

	if isNumber(r) {
		return intState
	}

	if isLetter(r) {
		return letterState
	}

	if r == '=' || r == '!' {
		l.next()             // consumme r
		if l.peek() == '=' { // next is '=' ?
			str := fmt.Sprintf("%s%s", string(r), string(l.peek()))
			// fmt.Printf("EQ or NEQ ? %q\n", str)
			if str == "!=" {
				// fmt.Println("Is NEQ")
				l.next()
				l.emit(token.NEQ)

				return parseNextState
			} else if str == "==" {
				//fmt.Println("Is EQ")
				l.next()
				l.emit(token.EQ)

				return parseNextState
			}
		} else {
			l.backup()
		}
	}

	if _, ok := symbols[r]; ok {
		return symbolState
	}

	return l.errorf("Don't know what to do with %q at line %d at %d.", r, l.linePos, l.charPos)
}

func spaceState(l *lexer) stateFn {
	// fmt.Println(">> spaceState")
	l.accept(" \t")
	l.startPos = l.currentPos
	return parseNextState
}

func newlineState(l *lexer) stateFn {
	// fmt.Println(">> newlineState")
	l.accept("\n")
	l.startPos = l.currentPos
	return parseNextState
}
func letterState(l *lexer) stateFn {
	// fmt.Println(">> letterState")

	var str string
	// consumme everything until the next space
	l.accept(identifier)
	str = l.input[l.startPos:l.currentPos]
	typ, ok := keywords[str]
	if ok {
		l.emit(typ)
		return parseNextState
	}

	//fmt.Printf(">> letterState :: identifier %s found", str)
	l.emit(token.IDENT)
	return parseNextState
}

func symbolState(l *lexer) stateFn {
	// fmt.Println(">> symbolState")
	c := l.next()

	typ, ok := symbols[c]
	if !ok {
		return l.errorf("Don't know what to do with %q at line %d at %d.", c, l.linePos, l.charPos)
	}

	l.emit(typ)
	return parseNextState
}

func intState(l *lexer) stateFn {
	// fmt.Println(">> intState")

	l.accept(digits)
	l.emit(token.INT)

	return parseNextState
}
