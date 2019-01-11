package main

import (
	"os"

	"github.com/tomocy/kinako/repl"
)

func main() {
	repl := repl.New(os.Stdin, os.Stdout)
	repl.Start()
}
