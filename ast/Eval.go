package ast

import "github.com/rogeralsing/GoMath/engine"

//Eval evaluates the result of the binary operation
func (node *BinaryNode) Eval(context *engine.Context) float64 {
	return node.Operator.Apply(node.Left, node.Right, context)
}

//Eval evaluates the value of the literal node
func (node *LiteralNode) Eval(context *engine.Context) float64 {
	return node.Value
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.Variable)
}
