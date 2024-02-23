package main

import (
	"fmt"
	"os"

	"github.com/pxlctzn/monkey/internal/lexer"
	"github.com/pxlctzn/monkey/internal/repl"
	"github.com/pxlctzn/monkey/internal/token"
	"golang.org/x/term"
)

const PROMPT = ">> "

func main_debug() {
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
		5 != 10;`

	var parsed []token.Token

	ch := lexer.New(input)
	for tk := range ch {
		parsed = append(parsed, tk)
	}

	for _, tk := range parsed {
		fmt.Printf("%s\n", tk.Debug())
	}

}
func main() {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		panic("stdin not found.")
	}

	previousState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), previousState)

	repl.Start(os.Stdin, os.Stdout)
}
