package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tomocy/kinako/evaluator"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/parser"
)

const prompt = "> "

type REPL struct {
	Reader io.Reader
	Writer io.Writer
}

func (r REPL) Start() {
	r.printPrompt()
	scanner := bufio.NewScanner(r.Reader)
	for scanner.Scan() {
		input := scanner.Text()
		r.printResult(input)
		r.printPrompt()
	}
}

func (r REPL) printPrompt() {
	fmt.Fprint(r.Writer, prompt)
}

func (r REPL) printResult(input string) {
	parser := parser.New(lexer.New(input))
	program := parser.ParseProgram()
	evaluator := evaluator.New(program)
	result := evaluator.Evaluate()
	fmt.Fprintln(r.Writer, result)
}
