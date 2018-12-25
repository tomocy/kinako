package evaluator

import (
	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/object"
)

type Evaluator struct {
	node ast.Node
}

func New(node ast.Node) *Evaluator {
	return &Evaluator{
		node: node,
	}
}

func (e *Evaluator) Evaluate() object.Object {
	switch e.node.(type) {
	case *ast.Integer:
		return e.evaluateInteger()
	default:
		return nil
	}
}

func (e *Evaluator) evaluateInteger() *object.Integer {
	integer := e.node.(*ast.Integer)
	return &object.Integer{
		Value: integer.Value,
	}
}
