package ast

import "github.com/rogeralsing/GoMath/engine"

func (node *BinaryNode) Eval(context *engine.Context) float64 {
	left := node.Left.Eval(context)
	switch {
	case OpAdd == node.Operator:
		return left + node.Right.Eval(context)
	case OpSub == node.Operator:
		return left - node.Right.Eval(context)
	case OpMul == node.Operator:
		return left * node.Right.Eval(context)
	case OpDiv == node.Operator:
		return left / node.Right.Eval(context)
	default:
		panic("unknown")
	}
}

func (node *LiteralNode) Eval(context *engine.Context) float64 {
	return node.Value
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.Variable)
}
