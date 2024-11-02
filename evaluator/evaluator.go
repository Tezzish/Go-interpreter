package evaluator

import (
	"Go-interpreter/ast"
	"Go-interpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBoolObject(node.Value)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node.Operator, Eval(node.Right))
	case *ast.InfixExpression:
		return evalInfixExpression(node.Operator, Eval(node.Left), Eval(node.Right))
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	}
	return NULL
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBoolObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalNegOperator(right)
	default:
		return nil
	}
}

func evalBangOperator(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalNegOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: 0 - value}
}

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBoolObject(left == right)
	case operator == "!=":
		return nativeBoolToBoolObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	if left.Type() != object.INTEGER_OBJ || right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	valueLeft := left.(*object.Integer).Value
	valueRight := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: valueLeft + valueRight}
	case "-":
		return &object.Integer{Value: valueLeft - valueRight}
	case "*":
		return &object.Integer{Value: valueLeft * valueRight}
	case "/":
		return &object.Integer{Value: valueLeft / valueRight}
	case "%":
		return &object.Integer{Value: valueLeft % valueRight}
	case "==":
		return nativeBoolToBoolObject(valueLeft == valueRight)
	case "<":
		return nativeBoolToBoolObject(valueLeft < valueRight)
	case ">":
		return nativeBoolToBoolObject(valueLeft > valueRight)
	case "!=":
		return nativeBoolToBoolObject(valueLeft != valueRight)
	default:
		return NULL
	}
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
	if isTruthy(condition) {
		return Eval(ie.Then)
	} else if ie.Else != nil {
		return Eval(ie.Else)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case FALSE:
		return false
	default:
		return true
	}

}
