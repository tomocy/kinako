package lexer

import (
	"testing"

	"github.com/tomocy/kinako/token"
)

func TestReadNextToken(t *testing.T) {
	input := "1"
	expects := []token.Token{
		{token.Integer, "1"},
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
