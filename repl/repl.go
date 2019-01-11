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
	reader    io.Reader
	writer    io.Writer
	evaluator *evaluator.Evaluator
}

func New(r io.Reader, w io.Writer) *REPL {
	return &REPL{
		reader:    r,
		writer:    w,
		evaluator: evaluator.New(),
	}
}

func (r REPL) Start() {
	r.printPrompt()
	scanner := bufio.NewScanner(r.reader)
	for scanner.Scan() {
		input := scanner.Text()
		r.printResult(input)
		r.printPrompt()
	}
}

func (r REPL) printPrompt() {
	fmt.Fprint(r.writer, prompt)
}

func (r REPL) printResult(input string) {
	parser := parser.New(lexer.New(input))
	program := parser.ParseProgram()
	result := r.evaluator.Evaluate(program)
	fmt.Fprintln(r.writer, result)
}
