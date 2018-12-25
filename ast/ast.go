package ast

import (
	"github.com/tomocy/kinako/token"
)

type Expression interface {
	expression()
}

type Integer struct {
	Token token.Token
	Value int64
}

func (i Integer) expression() {
}
