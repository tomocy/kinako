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
	return e.evaluate(e.node)
}

func (e *Evaluator) evaluate(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return e.evaluateProgram(node)
	case *ast.ExpressionStatement:
		return e.evaluateExpressionStatement(node)
	case *ast.Integer:
		return e.evaluateInteger(node)
	default:
		return nil
	}
}

func (e *Evaluator) evaluateProgram(node *ast.Program) object.Object {
	var obj object.Object
	for _, stmt := range node.Statements {
		obj = e.evaluate(stmt)
	}

	return obj
}

func (e *Evaluator) evaluateExpressionStatement(node *ast.ExpressionStatement) object.Object {
	return e.evaluate(node.Expression)
}

func (e *Evaluator) evaluateInteger(node *ast.Integer) *object.Integer {
	return &object.Integer{
		Value: node.Value,
	}
}
