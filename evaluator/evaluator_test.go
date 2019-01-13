package evaluator

import (
	"testing"

	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/object"
	"github.com/tomocy/kinako/parser"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
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
				Value: 15,
			},
		},
		{
			"(12 + 13) / 14;",
			&object.Integer{
				Value: 1,
			},
		},
		{
			"var x int;",
			&object.Integer{
				Value: 0,
			},
		},
		{
			"var x int = 16;",
			&object.Integer{
				Value: 16,
			},
		},
		{
			"var x int = 17; x;",
			&object.Integer{
				Value: 17,
			},
		},
		{
			"true;",
			&object.Boolean{
				Value: true,
			},
		},
		{
			"0; 0",
			&object.Error{
				Message: "failed to find semicolon",
			},
		},
		{
			"5 / 0;",
			&object.Error{
				Message: "divided by zero",
			},
		},
		{
			"y;",
			&object.Error{
				Message: "undefined variable: y",
			},
		},
	}
	for _, test := range tests {
		parser := parser.New(lexer.New(test.input))
		program := parser.ParseProgram()
		obj := New().Evaluate(program)
		switch obj := obj.(type) {
		case *object.Integer:
			testEvaluateInteger(t, obj, test.expected.(*object.Integer))
		case *object.Boolean:
			testEvaluateBoolean(t, obj, test.expected.(*object.Boolean))
		case *object.Error:
			testEvaluateError(t, obj, test.expected.(*object.Error))
		default:
			t.Fatalf("failed to assert type of object: %T, did you forget to add the type in switch?\n", obj)
		}
	}
}

func testEvaluateInteger(t *testing.T, actual, expected *object.Integer) {
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %d, but expected %d\n", actual.Value, expected.Value)
	}
}

func testEvaluateBoolean(t *testing.T, actual, expected *object.Boolean) {
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %t, but expected %t\n", actual.Value, expected.Value)
	}
}

func testEvaluateError(t *testing.T, actual, expected *object.Error) {
	if actual.Message != expected.Message {
		t.Errorf("unexpected message: got %s, but expected %s\n", actual.Message, expected.Message)
	}
}
