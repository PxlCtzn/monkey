package lexer

import "unicode"

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func isNumber(r rune) bool {
	return unicode.IsNumber(r)
}
