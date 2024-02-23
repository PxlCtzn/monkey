package token

type Type int

const (
	ILLEGAL Type = iota
	EOF
	// Identifiers and literals
	IDENT
	INT
	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	EQ
	NEQ
	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var strings = []string{
	"ILLEGAL",
	"EOF",
	"IDENT",
	"INT",
	"ASSIGN",
	"PLUS",
	"MINUS",
	"BANG",
	"ASTERISK",
	"SLASH",
	"LT",
	"GT",
	"EQ",
	"NEQ",
	"COMMA",
	"SEMICOLON",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"FUNCTION",
	"LET",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",
}

func (t Type) String() string {
	return strings[int(t)]
}
