package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/lexer"
	"github.com/tomocy/kinako/token"
)

var (
	ErrNoToken = errors.New("failed to find desired token")
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
	badStatements []*ast.BadStatement
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
		token.Minus:      p.parsePrefixExpression,
		token.Not:        p.parsePrefixExpression,
		token.LParen:     p.parseGroupExpression,
		token.Identifier: p.parseIdentifier,
		token.Integer:    p.parseInteger,
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
	for !p.has(token.EOF) {
		if stmt := p.parseStatement(); stmt != nil {
			stmts = append(stmts, stmt)
		}
		p.moveTokenForward()
	}

	return stmts
}

func (p *Parser) parseStatement() ast.Statement {
	var stmt ast.Statement
	switch p.currentToken.Type {
	case token.Var:
		stmt = p.parseVariableDeclaration()
	default:
		stmt = p.parseExpressionStatement()
	}

	if err := p.expectAndMoveTokenForward(token.Semicolon); err != nil {
		p.keepBadStatement("failed to find semicolon")
	}

	if p.hasBadStatements() {
		return p.takeOutLastBadStatement()
	}

	return stmt
}

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	if err := p.expectAndMoveTokenForward(token.Identifier); err != nil {
		p.keepBadStatement("failed to find identifier of variable")
		return nil
	}
	stmt := &ast.VariableDeclaration{
		Identifier: p.parseIdentifier().(*ast.Identifier),
	}

	if err := p.expectAndMoveTokenForward(token.Identifier); err != nil {
		p.keepBadStatement("failed to find type name of variable")
		return nil
	}
	stmt.Type = p.parseIdentifier().(*ast.Identifier)

	if err := p.expectAndMoveTokenForward(token.Assign); err != nil {
		return stmt
	}

	p.moveTokenForward()
	stmt.Expression = p.parseExpression(lowest)

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Expression: p.parseExpression(lowest),
	}
}

func (p *Parser) parseExpression(prio priority) ast.Expression {
	expr := p.prefixParsers[p.currentToken.Type]()
	for !p.willHave(token.Semicolon) && p.checkReadingTokenPriority().isHigherThan(prio) {
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
	if err := p.expectAndMoveTokenForward(token.RParen); err != nil {
		p.keepBadStatement("failed to find rparen")
		return nil
	}

	return expr
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Name: p.currentToken.Literal,
	}
}

func (p *Parser) parseInteger() ast.Expression {
	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("faild to parse %s to integer: %s", p.currentToken.Literal, err))
	}

	return &ast.Integer{
		Value: value,
	}
}

func (p *Parser) hasBadStatements() bool {
	return 0 < len(p.badStatements)
}

func (p *Parser) takeOutLastBadStatement() *ast.BadStatement {
	i := len(p.badStatements) - 1
	stmt := p.badStatements[i]
	p.badStatements = p.badStatements[:i]
	return stmt
}

func (p *Parser) keepBadStatement(msg string) {
	p.badStatements = append(p.badStatements, &ast.BadStatement{
		Message: msg,
	})
}

func (p *Parser) reportBadStatement(msg string) *ast.BadStatement {
	return &ast.BadStatement{
		Message: msg,
	}
}

func (p *Parser) moveFirstTwoTokenForward() {
	p.moveTokenForward()
	p.moveTokenForward()
}

func (p *Parser) expectAndMoveTokenForward(t token.Type) error {
	if !p.willHave(t) {
		return ErrNoToken
	}
	p.moveTokenForward()

	return nil
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

func (p Parser) has(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p Parser) willHave(t token.Type) bool {
	return p.readingToken.Type == t
}
