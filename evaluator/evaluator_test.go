package evaluator

import (
	"testing"

	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/object"
	"github.com/tomocy/kinako/parser"
)

func TestEvaluateInteger(t *testing.T) {
	input := "5"
	expect := &object.Integer{
		Value: 5,
	}
	parser := parser.New(lexer.New(input))
	program := parser.ParseProgram()
	evaluator := New(program)
	obj := evaluator.Evaluate()
	integer, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("unexpected object: got %T, expected *object.Integer", obj)
	}
	if integer.Value != expect.Value {
		t.Errorf("unexpected value: got %d, expected %d\n", integer.Value, expect.Value)
	}
}
