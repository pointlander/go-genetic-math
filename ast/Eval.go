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

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.Variable)
}
