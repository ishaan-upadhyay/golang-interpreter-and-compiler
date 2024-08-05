package repl
/*
 * The REPL (read-eval-print loop) for Monkey.
 */

import (
	"bufio"
	"fmt"
	"io"
	"golang-interpreter-and-compiler/src/monkey/lexer"
	"golang-interpreter-and-compiler/src/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text();
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}