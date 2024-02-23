package repl

import (
	"fmt"
	"io"
	"os/user"

	"github.com/pxlctzn/monkey/internal/lexer"
	"golang.org/x/term"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	ioWR := struct {
		io.Reader
		io.Writer
	}{
		in,
		out,
	}

	tty := term.NewTerminal(ioWR, PROMPT)

	{ // Prompt welcome message
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		tty.Write(bytesf("Hello %q! This is the Monkey Programming Language !\n", user.Username))
		tty.Write(bytesf("Fell free to type in commands\n\n"))
	}

	for {
		line, err := tty.ReadLine()
		if err == io.EOF {
			tty.Write(bytesf("Okay! Byyyyee\n"))
			return
		}
		ch := lexer.New(line)
		for tk := range ch {
			tty.Write(bytesf("%v\n", tk.Debug()))
		}
	}
}

func bytesf(format string, v ...any) []byte {
	return []byte(fmt.Sprintf(format, v...))
}
