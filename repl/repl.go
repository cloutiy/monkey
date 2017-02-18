// repl/repl.go
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">>"

//Starts a REPL
func Start(in io.Reader, out io.Writer) {

	//Create a new string scanner thatreads from stdin
	scanner := bufio.NewScanner(in)

	//Infinitely
	for {
		//print the REPL prompt
		fmt.Printf(PROMPT)

		//get a line of input from the terminal
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		//Copy the buffer into a string
		line := scanner.Text()

		//Initialize a new lexer with this string
		l := lexer.New(line)

		//Print the results of lexing until we get to EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok := l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
