// main.go
package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s; Welcome to the future.\n", user.Username)
	fmt.Printf("Feel free to type in commands.")
	repl.Start(os.Stdin, os.Stdout)

}
