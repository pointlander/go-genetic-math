package ast

import "github.com/rogeralsing/GoMath/engine"

func (node *AddNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) + node.Right.Eval(context)
}

func (node *DivNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) / node.Right.Eval(context)
}

func (node *LiteralNode) Eval(context *engine.Context) float64 {
	return node.Value
}

func (node *MulNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) * node.Right.Eval(context)
}

func (node *SubNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) - node.Right.Eval(context)
}

func (node *BinaryNode) Eval(context *engine.Context) float64 {
	left := node.Left.Eval(context)
	right := node.Right.Eval(context)

	switch {
	case AddOp == node.Operator:
		return left + right
	case SubOp == node.Operator:
		return left - right
	case MulOp == node.Operator:
		return left * right
	case DivOp == node.Operator:
		return left / right
	}
	return 0
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.Variable)
}
