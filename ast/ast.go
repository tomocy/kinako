package ast

import (
	"github.com/tomocy/kinako/token"
)

type Operator string

const (
	Minus Operator = "-"
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

type Expression interface {
	Node
	expression()
}

type PrefixExpression struct {
	Operator    Operator
	RExpression Expression
}

func (e PrefixExpression) node() {
}

func (e PrefixExpression) expression() {
}

type Integer struct {
	Token token.Token
	Value int64
}

func (i Integer) node() {
}

func (i Integer) expression() {
}
