package evaluator

import (
	"testing"

	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/object"
	"github.com/tomocy/kinako/parser"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		input  string
		expect object.Object
	}{
		{
			"5;",
			&object.Integer{
				Value: 5,
			},
		},
		{
			"-6;",
			&object.Integer{
				Value: -6,
			},
		},
		{
			"7 + 8 - 9 * 10 / 11;",
			&object.Integer{
				Value: 5,
			},
		},
	}
	for _, test := range tests {
		parser := parser.New(lexer.New(test.input))
		program := parser.ParseProgram()
		evaluator := New(program)
		obj := evaluator.Evaluate()
		switch obj := obj.(type) {
		case *object.Integer:
			testEvaluateIntegerObject(t, obj, test.expect.(*object.Integer))
		default:
			t.Fatalf("failed to assert type of object: %T, did you forget to add the type in switch?\n", obj)
		}
	}
}

func testEvaluateIntegerObject(t *testing.T, actual, expected *object.Integer) {
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %d, but expected %d\n", actual.Value, expected.Value)
	}
}
