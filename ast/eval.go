package ast

import "github.com/rogeralsing/go-genetic-math/engine"

//Eval evaluates the result of the binary operation
func (node *BinaryNode) Eval(context *engine.Context) float64 {
	return node.operator.Apply(node.left, node.right, context)
}

//Eval evaluates the value of the literal node
func (node *LiteralNode) Eval(context *engine.Context) float64 {
	return node.value
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.variable)
}
