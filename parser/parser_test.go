package parser

import (
	"log"
	"testing"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/token"
)

func TestParseProgram(t *testing.T) {
	input := `
	5
	-6
	`
	expecteds := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.Integer{
				Token: token.Token{
					Type:    token.Integer,
					Literal: "5",
				},
				Value: 5,
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator: ast.Minus,
				RExpression: &ast.Integer{
					Token: token.Token{
						Type:    token.Integer,
						Literal: "6",
					},
					Value: 6,
				},
			},
		},
	}
	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	for i := 0; i < len(expecteds); i++ {
		log.Println(i)
		testParseStatement(t, program.Statements[i], expecteds[i])
		log.Println(i)
	}
}

func testParseStatement(t *testing.T, actual, expected ast.Statement) {
	switch actual := actual.(type) {
	case *ast.ExpressionStatement:
		testParseExpressionStatement(t, actual, expected.(*ast.ExpressionStatement))
	default:
		t.Fatalf("unexpected type of statement: %T\n", actual)
	}
}

func testParseExpressionStatement(t *testing.T, actual, expected *ast.ExpressionStatement) {
	testParseExpression(t, actual.Expression, expected.Expression)
}

func testParseExpression(t *testing.T, actual, expected ast.Expression) {
	switch actual := actual.(type) {
	case *ast.PrefixExpression:
		testParsePrefixExpression(t, actual, expected.(*ast.PrefixExpression))
	case *ast.Integer:
		testParseInteger(t, actual, expected.(*ast.Integer))
	default:
		t.Fatalf("unexpected type of expression: %T\n", actual)
	}
}

func testParsePrefixExpression(t *testing.T, actual, expected *ast.PrefixExpression) {
	if actual.Operator != expected.Operator {
		t.Errorf("unexpected operator: got %s, but expected %s\n", actual.Operator, expected.Operator)
	}

	testParseExpression(t, actual.RExpression, expected.RExpression)
}

func testParseInteger(t *testing.T, actual, expected *ast.Integer) {
	if actual.Token != expected.Token {
		t.Errorf("unexpected token: got %v, but expected %v\n", actual.Token, expected.Token)
	}
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %d, but expected %d\n", actual.Value, expected.Value)
	}
}
