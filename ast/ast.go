package ast

import (
	"github.com/tomocy/kinako/token"
)

type Node interface {
	node()
}

type Program struct {
	Statements []Statement
}

func (p Program) node() {
}

type Statement interface {
	Node
	statement()
}

type ExpressionStatement struct {
	Expression Expression
}

func (s ExpressionStatement) node() {
}

func (s ExpressionStatement) statement() {
}

type VariableDeclaration struct {
	Identifier string
	Type       string
	Expression Expression
}

func (s VariableDeclaration) node() {
}

func (s VariableDeclaration) statement() {
}

type BadStatement struct {
	Message string
}

func (s BadStatement) node() {
}

func (s BadStatement) statement() {
}

type Expression interface {
	Node
	expression()
}

type PrefixExpression struct {
	Operator    PrefixOperator
	RExpression Expression
}

type PrefixOperator string

const (
	Negative PrefixOperator = "-"
)

var PrefixOperators = map[token.Type]PrefixOperator{
	token.Minus: Negative,
}

func (e PrefixExpression) node() {
}

func (e PrefixExpression) expression() {
}

type InfixOperator string

const (
	Plus     InfixOperator = "+"
	Minus                  = "-"
	Asterisk               = "*"
	Slash                  = "/"
)

var InfixOperators = map[token.Type]InfixOperator{
	token.Plus:     Plus,
	token.Minus:    Minus,
	token.Asterisk: Asterisk,
	token.Slash:    Slash,
}

type InfixExpression struct {
	LExpression Expression
	Operator    InfixOperator
	RExpression Expression
}

func (e InfixExpression) node() {
}

func (e InfixExpression) expression() {
}

type Integer struct {
	Token token.Token
	Value int64
}

func (e Integer) node() {
}

func (e Integer) expression() {
}
