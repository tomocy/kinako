package evaluator

import (
	"fmt"

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
	case *ast.BadStatement:
		return e.evaluateBadStatement(node)
	case *ast.PrefixExpression:
		return e.evaluatePrefixExpression(node)
	case *ast.InfixExpression:
		return e.evaluateInfixExpression(node)
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

func (e *Evaluator) evaluateBadStatement(node *ast.BadStatement) object.Object {
	return &object.Error{
		Message: node.Message,
	}
}

func (e *Evaluator) evaluatePrefixExpression(node *ast.PrefixExpression) object.Object {
	obj := e.evaluate(node.RExpression)
	switch node.Operator {
	case ast.Negative:
		return e.evaluateNegativeInteger(obj.(*object.Integer))
	default:
		panic(fmt.Sprintf("failed to assert prefix operator type because of developer. contact him or her to inform the missing type is %s", node.Operator))
	}
}

func (e *Evaluator) evaluateNegativeInteger(obj *object.Integer) *object.Integer {
	return &object.Integer{
		Value: -1 * obj.Value,
	}
}

func (e *Evaluator) evaluateInfixExpression(node *ast.InfixExpression) object.Object {
	left := e.evaluate(node.LExpression)
	right := e.evaluate(node.RExpression)
	switch node.Operator {
	case ast.Plus:
		return e.evaluateAddition(left, right)
	case ast.Minus:
		return e.evaluateSubtraction(left, right)
	case ast.Asterisk:
		return e.evaluateMultiplication(left, right)
	case ast.Slash:
		return e.evaluateDivision(left, right)
	default:
		panic(fmt.Sprintf("failed to assert infix operator type because of developer. contact him or her to inform the missing type is %s", node.Operator))
	}
}

func (e *Evaluator) evaluateAddition(left, right object.Object) object.Object {
	return &object.Integer{
		Value: left.(*object.Integer).Value + right.(*object.Integer).Value,
	}
}

func (e *Evaluator) evaluateSubtraction(left, right object.Object) object.Object {
	return &object.Integer{
		Value: left.(*object.Integer).Value - right.(*object.Integer).Value,
	}
}

func (e *Evaluator) evaluateMultiplication(left, right object.Object) object.Object {
	return &object.Integer{
		Value: left.(*object.Integer).Value * right.(*object.Integer).Value,
	}
}

func (e *Evaluator) evaluateDivision(left, right object.Object) object.Object {
	rightVal := right.(*object.Integer).Value
	if rightVal == 0 {
		panic("divided by zero")
	}

	return &object.Integer{
		Value: left.(*object.Integer).Value / rightVal,
	}
}

func (e *Evaluator) evaluateInteger(node *ast.Integer) *object.Integer {
	return &object.Integer{
		Value: node.Value,
	}
}
