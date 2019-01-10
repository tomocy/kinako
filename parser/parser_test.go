package parser

import (
	"testing"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
)

func TestParseProgram(t *testing.T) {
	input := `
	5; -6;
	7 + 8 - 9 * 10 / 11;
	(12 + 13) / 14;
	15; 16
	`
	expecteds := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.Integer{
				Value: 5,
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator: ast.Negative,
				RExpression: &ast.Integer{
					Value: 6,
				},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				LExpression: &ast.Integer{
					Value: 7,
				},
				Operator: ast.Plus,
				RExpression: &ast.InfixExpression{
					LExpression: &ast.Integer{
						Value: 8,
					},
					Operator: ast.Minus,
					RExpression: &ast.InfixExpression{
						LExpression: &ast.Integer{
							Value: 9,
						},
						Operator: ast.Asterisk,
						RExpression: &ast.InfixExpression{
							LExpression: &ast.Integer{
								Value: 10,
							},
							Operator: ast.Slash,
							RExpression: &ast.Integer{
								Value: 11,
							},
						},
					},
				},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				LExpression: &ast.InfixExpression{
					LExpression: &ast.Integer{
						Value: 12,
					},
					Operator: ast.Plus,
					RExpression: &ast.Integer{
						Value: 13,
					},
				},
				Operator: ast.Slash,
				RExpression: &ast.Integer{
					Value: 14,
				},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.Integer{
				Value: 15,
			},
		},
		&ast.BadStatement{
			Message: "failed to find semicolon. semicolon should be at the end of a statement",
		},
	}
	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	for i := 0; i < len(expecteds); i++ {
		testParseStatement(t, program.Statements[i], expecteds[i])
	}
}

func testParseStatement(t *testing.T, actual, expected ast.Statement) {
	switch actual := actual.(type) {
	case *ast.ExpressionStatement:
		testParseExpressionStatement(t, actual, expected.(*ast.ExpressionStatement))
	case *ast.BadStatement:
		testParseBadStatement(t, actual, expected.(*ast.BadStatement))
	default:
		t.Fatalf("failed to assert type of statement: %T, did you forget to add the type in switch?\n", actual)
	}
}

func testParseExpressionStatement(t *testing.T, actual, expected *ast.ExpressionStatement) {
	testParseExpression(t, actual.Expression, expected.Expression)
}

func testParseExpression(t *testing.T, actual, expected ast.Expression) {
	switch actual := actual.(type) {
	case *ast.PrefixExpression:
		testParsePrefixExpression(t, actual, expected.(*ast.PrefixExpression))
	case *ast.InfixExpression:
		testParseInfixExpression(t, actual, expected.(*ast.InfixExpression))
	case *ast.Integer:
		testParseInteger(t, actual, expected.(*ast.Integer))
	default:
		t.Fatalf("failed to assert type of expression: %T, did you forget to add the type in switch?\n", actual)
	}
}

func testParsePrefixExpression(t *testing.T, actual, expected *ast.PrefixExpression) {
	if actual.Operator != expected.Operator {
		t.Errorf("unexpected operator: got %s, but expected %s\n", actual.Operator, expected.Operator)
	}
	testParseExpression(t, actual.RExpression, expected.RExpression)
}

func testParseInfixExpression(t *testing.T, actual, expected *ast.InfixExpression) {
	testParseExpression(t, actual.LExpression, expected.LExpression)
	if actual.Operator != expected.Operator {
		t.Errorf("unexpected infix operator: got %s, but expected %s\n", actual.Operator, expected.Operator)
	}
	testParseExpression(t, actual.RExpression, expected.RExpression)
}

func testParseInteger(t *testing.T, actual, expected *ast.Integer) {
	if actual.Value != expected.Value {
		t.Errorf("unexpected value: got %d, but expected %d\n", actual.Value, expected.Value)
	}
}

func testParseBadStatement(t *testing.T, actual, expected *ast.BadStatement) {
	if actual.Message != expected.Message {
		t.Errorf("unexpected message: got %s, but expected %s\n", actual.Message, expected.Message)
	}
}
