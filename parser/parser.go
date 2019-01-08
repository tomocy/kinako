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

func (p *Parser) ParseProgram() *ast.Program {
	return &ast.Program{
		Statements: p.parseStatements(),
	}
}

func (p *Parser) parseStatements() []ast.Statement {
	stmts := make([]ast.Statement, 0)
	for !p.hasEOF() {
		if stmt := p.parseStatement(); stmt != nil {
			stmts = append(stmts, stmt)
		}
		p.readToken()
	}

	return stmts
}

func (p *Parser) parseStatement() ast.Statement {
	return p.parseExpressionStatement()
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Expression: p.parseExpression(),
	}
}

func (p *Parser) parseExpression() ast.Expression {
	if p.has(token.Minus) {
		return p.parsePrefixExpression()
	}

	return p.parseInteger()
}

func (p *Parser) parsePrefixExpression() *ast.PrefixExpression {
	expr := &ast.PrefixExpression{
		Operator: ast.Operator(p.currentToken.Literal),
	}
	p.readToken()
	expr.RExpression = p.parseExpression()

	return expr
}

func (p *Parser) parseInteger() *ast.Integer {
	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("faild to parse integer: %s", err))
	}

	return &ast.Integer{
		Token: p.currentToken,
		Value: value,
	}
}

func (p *Parser) readToken() {
	p.currentToken = p.lexer.ReadNextToken()
}

func (p Parser) hasEOF() bool {
	return p.has(token.EOF)
}

func (p Parser) has(t token.Type) bool {
	return p.currentToken.Type == t
}
