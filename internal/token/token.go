package token

import "fmt"

type Token struct {
	typ  Type
	val  string
	line int
	pos  int
}

func (t Token) Type() Type {
	return t.typ
}
func (t Token) Is(typ Type) bool {
	return t.typ == typ
}

func (t Token) Debug() string {
	var val string
	if t.typ == EOF {
		val = "EOF"
	} else {
		val = t.val
	}
	return fmt.Sprintf("Line %d at %d <%s> %q", t.line, t.pos, t.typ, val)
}

func (t Token) String() string {
	if t.typ == EOF {
		return "EOF"
	}
	return t.val
}

func New(t Type, v string, l int, p int) Token {
	tk := Token{t, v, l, p}
	return tk
}
