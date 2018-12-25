package parser

import (
	"fmt"
	"strconv"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.readToken()

	return p
}

func (p *Parser) readToken() {
	p.currentToken = p.lexer.ReadNextToken()
}

func (p *Parser) ParseExpression() ast.Expression {
	return p.parserInteger()
}

func (p *Parser) parserInteger() *ast.Integer {
	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("faild to parse integer: %s", err))
	}

	return &ast.Integer{
		Token: p.currentToken,
		Value: value,
	}
}
