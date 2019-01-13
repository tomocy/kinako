package evaluator

import (
	"fmt"

	"github.com/tomocy/kinako/ast"
	"github.com/tomocy/kinako/object"
)

var zeroValues = map[string]object.Object{
	"int": &object.Integer{
		Value: 0,
	},
}

type Evaluator struct {
	env Environment
}

func New() *Evaluator {
	return &Evaluator{
		env: NewEnvironment(),
	}
}

func (e *Evaluator) Evaluate(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return e.evaluateProgram(node)
	case *ast.ExpressionStatement:
		return e.evaluateExpressionStatement(node)
	case *ast.VariableDeclaration:
		return e.evaluateVariableDeclaration(node)
	case *ast.BadStatement:
		return e.evaluateBadStatement(node)
	case *ast.PrefixExpression:
		return e.evaluatePrefixExpression(node)
	case *ast.InfixExpression:
		return e.evaluateInfixExpression(node)
	case *ast.Identifier:
		return e.evaluateIdentifier(node)
	case *ast.Integer:
		return e.evaluateInteger(node)
	default:
		return nil
	}
}

func (e *Evaluator) evaluateProgram(node *ast.Program) object.Object {
	var obj object.Object
	for _, stmt := range node.Statements {
		obj = e.Evaluate(stmt)
	}

	return obj
}

func (e *Evaluator) evaluateExpressionStatement(node *ast.ExpressionStatement) object.Object {
	return e.Evaluate(node.Expression)
}

func (e *Evaluator) evaluateVariableDeclaration(node *ast.VariableDeclaration) object.Object {
	var obj object.Object
	if node.Expression == nil {
		obj = zeroValues[node.Type.Name]
	} else {
		obj = e.Evaluate(node.Expression)
	}

	if err := e.env.Set(node.Identifier.Name, obj); err != nil {
		return &object.Error{
			Message: err.Error(),
		}
	}

	return obj
}

func (e *Evaluator) evaluateBadStatement(node *ast.BadStatement) object.Object {
	return &object.Error{
		Message: node.Message,
	}
}

func (e *Evaluator) evaluatePrefixExpression(node *ast.PrefixExpression) object.Object {
	obj := e.Evaluate(node.RExpression)
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
	left := e.Evaluate(node.LExpression)
	right := e.Evaluate(node.RExpression)
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
		return &object.Error{
			Message: "divided by zero",
		}
	}

	return &object.Integer{
		Value: left.(*object.Integer).Value / rightVal,
	}
}

func (e *Evaluator) evaluateIdentifier(node *ast.Identifier) object.Object {
	if obj, ok := e.env[node.Name]; ok {
		return obj
	}

	return &object.Error{
		Message: fmt.Sprintf("undefined variable: %s", node.Name),
	}
}

func (e *Evaluator) evaluateInteger(node *ast.Integer) *object.Integer {
	return &object.Integer{
		Value: node.Value,
	}
}
