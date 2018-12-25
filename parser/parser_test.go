package parser

import (
	"testing"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
)

func TestParseExpression(t *testing.T) {
	input := `
	5
	`
	expect := &ast.Integer{
		Value: 5,
	}
	parser := New(lexer.New(input))
	expr := parser.ParseExpression()
	integer, ok := expr.(*ast.Integer)
	if !ok {
		t.Errorf("unexpected expression: got %T, expected *ast.Integer\n", expr)
	}
	if integer.Value != expect.Value {
		t.Errorf("unexpected value: got %d, expected %d\n", integer.Value, expect.Value)
	}
}
