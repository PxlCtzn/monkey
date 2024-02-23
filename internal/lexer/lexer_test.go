package lexer

import (
	"fmt"
	"testing"

	"github.com/pxlctzn/monkey/internal/token"
)

func TestLexer(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y){
		x + y;
	};
	
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	5 == 10;
	5 != 10;
	`

	tests := []struct {
		expectedType  token.Type
		expectedValue string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "5"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.NEQ, "!="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	var parsed []token.Token

	ch := New(input)
	for tk := range ch {
		parsed = append(parsed, tk)
	}

	for i, tk := range parsed {
		fmt.Printf("%s\n", tk.Debug())
		if !tk.Is(tests[i].expectedType) {
			t.Fatalf("Error Token Type.\nExpected: %q\nGot: %q\nDebug:\n%s\n", tests[i].expectedType, tk.Type(), tk.Debug())
		}
		if tk.String() != tests[i].expectedValue {
			t.Fatalf("Error Token Value.\nExpected: %q\nGot: %q\nDebug:\n%s\n", tests[i].expectedValue, tk.String(), tk.Debug())
		}
	}
}

func TestSinglCharToken(t *testing.T) {
	input := "=+{}(),;"

	tests := []struct {
		expectedType  token.Type
		expectedValue string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	var parsed []token.Token

	ch := New(input)
	for tk := range ch {
		parsed = append(parsed, tk)
	}

	for i, tk := range parsed {
		fmt.Printf("%s\n", tk.Debug())
		if !tk.Is(tests[i].expectedType) {
			t.Fatalf("Error Token Type.\nExpected: %q\nGot: %q\nDebug:\n%s\n", tests[i].expectedType, tk.Type(), tk.Debug())
		}
		if tk.String() != tests[i].expectedValue {
			t.Fatalf("Error Token Value.\nExpected: %q\nGot: %q\nDebug:\n%s\n", tests[i].expectedValue, tk.String(), tk.Debug())
		}
	}

}

func TestEqNeq(t *testing.T) {
	input := `10 != 20;
	5 == 10;`
	tests := []token.Type{
		token.INT,
		token.NEQ,
		token.INT,
		token.SEMICOLON,
		token.INT,
		token.EQ,
		token.INT,
		token.SEMICOLON,
	}

	var parsed []token.Token

	ch := New(input)
	for tk := range ch {
		parsed = append(parsed, tk)
	}

	for i, tk := range parsed {
		fmt.Printf("%s\n", tk.Debug())
		if !tk.Is(tests[i]) {
			t.Fatalf("Error Token Type.\nExpected: %q\nGot: %q\nDebug:\n%s\n", tests[i], tk.Type(), tk.Debug())
		}
	}
}
