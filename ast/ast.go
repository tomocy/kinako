package ast

import (
	"github.com/tomocy/kinako/token"
)

type Node interface {
	node()
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}

type Integer struct {
	Token token.Token
	Value int64
}

func (i Integer) node() {
}

func (i Integer) expression() {
}
