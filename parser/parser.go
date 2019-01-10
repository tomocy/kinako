package parser

import (
	"fmt"
	"strconv"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/token"
)

type priority int

const (
	_ priority = iota
	lowest
	additive
	multiplicative
	prefix
	infix
)

var precedence = map[token.Type]priority{
	token.Plus:     additive,
	token.Minus:    additive,
	token.Asterisk: multiplicative,
	token.Slash:    multiplicative,
}

func (p priority) isHigherThan(prec priority) bool {
	return p > prec
}

type prefixParser func() ast.Expression

type infixParser func(ast.Expression) ast.Expression

type Parser struct {
	lexer         *lexer.Lexer
	prefixParsers map[token.Type]prefixParser
	infixParsers  map[token.Type]infixParser
	currentToken  token.Token
	readingToken  token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.moveFirstTwoTokenForward()
	p.registerPrefixParsers()
	p.registerInfixParsers()

	return p
}

func (p *Parser) registerPrefixParsers() {
	p.prefixParsers = map[token.Type]prefixParser{
		token.Minus:   p.parsePrefixExpression,
		token.LParen:  p.parseGroupExpression,
		token.Integer: p.parseInteger,
	}
}

func (p *Parser) registerInfixParsers() {
	p.infixParsers = map[token.Type]infixParser{
		token.Plus:     p.parseInfixExpression,
		token.Minus:    p.parseInfixExpression,
		token.Asterisk: p.parseInfixExpression,
		token.Slash:    p.parseInfixExpression,
	}
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
		p.moveTokenForward()
	}

	return stmts
}

func (p *Parser) parseStatement() ast.Statement {
	stmt := p.parseExpressionStatement()
	if !p.willHaveSemicolon() {
		return p.parseBadStatement("failed to find semicolon. semicolon should be at the end of a statement")
	}
	p.moveTokenForward()

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Expression: p.parseExpression(lowest),
	}
}

func (p *Parser) parseExpression(prio priority) ast.Expression {
	expr := p.prefixParsers[p.currentToken.Type]()
	for !p.willHaveSemicolon() && p.checkReadingTokenPriority().isHigherThan(prio) {
		p.moveTokenForward()
		expr = p.infixParsers[p.currentToken.Type](expr)
	}

	return expr
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Operator: ast.PrefixOperators[p.currentToken.Type],
	}
	p.moveTokenForward()
	expr.RExpression = p.parseExpression(prefix)

	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		LExpression: left,
		Operator:    ast.InfixOperators[p.currentToken.Type],
	}
	p.moveTokenForward()
	expr.RExpression = p.parseExpression(p.checkCurrentTokenPriority())

	return expr
}

func (p *Parser) parseGroupExpression() ast.Expression {
	p.moveTokenForward()
	expr := p.parseExpression(lowest)
	if !p.willHave(token.RParen) {
		panic("failed to find rparen which is the pair of lparen")
	}
	p.moveTokenForward()

	return expr
}

func (p *Parser) parseInteger() ast.Expression {
	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("faild to parse %s to integer: %s", p.currentToken.Literal, err))
	}

	return &ast.Integer{
		Token: p.currentToken,
		Value: value,
	}
}

func (p *Parser) parseBadStatement(msg string) *ast.BadStatement {
	return &ast.BadStatement{
		Message: msg,
	}
}

func (p *Parser) moveFirstTwoTokenForward() {
	p.moveTokenForward()
	p.moveTokenForward()
}

func (p *Parser) moveTokenForward() {
	p.currentToken = p.readingToken
	p.readingToken = p.lexer.ReadNextToken()
}

func (p Parser) checkCurrentTokenPriority() priority {
	return checkTokenPriority(p.currentToken.Type)
}

func (p Parser) checkReadingTokenPriority() priority {
	return checkTokenPriority(p.readingToken.Type)
}

func checkTokenPriority(t token.Type) priority {
	if prio, ok := precedence[t]; ok {
		return prio
	}

	return lowest
}

func (p Parser) hasEOF() bool {
	return p.has(token.EOF)
}

func (p Parser) willHaveSemicolon() bool {
	return p.willHave(token.Semicolon)
}

func (p Parser) has(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p Parser) willHave(t token.Type) bool {
	return p.readingToken.Type == t
}
