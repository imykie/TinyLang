package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/imykie/tinylang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to TinyLang!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
