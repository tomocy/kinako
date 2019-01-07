package parser

import (
	"testing"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/token"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected ast.Expression
	}{
		{
			"5",
			&ast.Integer{
				Token: token.Token{
					Type:    token.Integer,
					Literal: "5",
				},
				Value: 5,
			},
		},
	}
	for _, test := range tests {
		parser := New(lexer.New(test.input))
		expr := parser.parseExpression()
		switch expr := expr.(type) {
		case *ast.Integer:
			testParseIntegerExpression(t, expr, test.expected.(*ast.Integer))
		default:
			t.Fatalf("unexpected type of expression: %T\n", expr)
		}
	}
}

func testParseIntegerExpression(t *testing.T, actual, expected *ast.Integer) {
	if actual.Token != expected.Token {
		t.Errorf("unexpected token: got %v, but expected %v\n", actual.Token, expected.Token)
	}
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %d, but expected %d\n", actual.Value, expected.Value)
	}
}
