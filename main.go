package main

import (
	"os"

	"github.com/tomocy/kinako/repl"
)

func main() {
	repl := repl.REPL{
		Reader: os.Stdin,
		Writer: os.Stdout,
	}
	repl.Start()
}
