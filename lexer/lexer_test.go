package lexer

import (
	"testing"

	"github.com/tomocy/kinako/token"
)

func TestReadNextToken(t *testing.T) {
	input := `
	1 + 2 - 3 * 4 / 5
	((6 + 7) * (8 - 9)) / 10
	`
	expects := []token.Token{
		{token.Integer, "1"}, {token.Plus, "+"}, {token.Integer, "2"}, {token.Minus, "-"}, {token.Integer, "3"}, {token.Asterisk, "*"}, {token.Integer, "4"}, {token.Slash, "/"}, {token.Integer, "5"},
		{token.LParen, "("},
		{token.LParen, "("}, {token.Integer, "6"}, {token.Plus, "+"}, {token.Integer, "7"}, {token.RParen, ")"},
		{token.Asterisk, "*"},
		{token.LParen, "("}, {token.Integer, "8"}, {token.Minus, "-"}, {token.Integer, "9"}, {token.RParen, ")"},
		{token.RParen, ")"}, {token.Slash, "/"}, {token.Integer, "10"},
		{token.EOF, ""},
	}
	lexer := New(input)
	for _, expect := range expects {
		token := lexer.ReadNextToken()
		if token.Type != expect.Type {
			t.Errorf("unexpected token type: got %v, but expected %v", token.Type, expect.Type)
		}
		if token.Literal != expect.Literal {
			t.Errorf("unexpected token literal: got %v, but expected %v", token.Literal, expect.Literal)
		}
	}
}
