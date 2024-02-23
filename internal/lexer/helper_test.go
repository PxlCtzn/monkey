package lexer

import (
	"fmt"
	"testing"
)

func TestIsSpace(t *testing.T) {
	input := `a 0
	,!`
	expected := []bool{
		false, // a
		true,  // space
		false, // 0
		false, // EOL
		true,  // tab
		false, // ,
		false, // !
	}

	for i, c := range input {
		if expected[i] != isSpace(c) {
			t.Fatalf("Expected isSpace(%q) to return %+v but got %+v.", c, expected[i], isSpace(c))
		}
	}
}

func TestIsEOL(t *testing.T) {
	input := `a 0
	,!`
	expected := []bool{
		false, // a
		false, // space
		false, // 0
		true,  // EOL
		false, // tab
		false, // ,
		false, // !
	}

	for i, c := range input {
		if expected[i] != isEndOfLine(c) {
			t.Fatalf("Expected isEndOfLine(%q) to return %+v but got %+v.", c, expected[i], isSpace(c))
		}
	}
}

func TestIsLetter(t *testing.T) {
	input := `a é0Zç
	,!`
	expected := []bool{
		true,  // a
		false, // space
		true,  // é
		false, // 0
		true,  // Z
		true,  // ç
		false, // EOL
		false, // tab
		false, // ,
		false, // !
	}
	i := 0

	for _, c := range input {
		if expected[i] != isLetter(c) {
			t.Fatalf("Expected isLetter(%q) to return %+v but got %+v.", c, expected[i], isLetter(c))
		}
		i++
	}
}

func TestIsNumber(t *testing.T) {
	input := `a é064Zç
	,!`
	expected := []bool{
		false, // a
		false, // space
		false, // é
		true,  // 0
		true,  // 6
		true,  // 4
		false, // Z
		false, // ç
		false, // EOL
		false, // tab
		false, // ,
		false, // !
	}

	i := 0
	for _, c := range input {
		fmt.Printf("%d %q ; expected: %+v\n", i, c, expected[i])
		if expected[i] != isNumber(c) {
			t.Fatalf("Expected isNumber(%q) to return %+v but got %+v.", c, expected[i], isNumber(c))
		}
		i++
	}
}
