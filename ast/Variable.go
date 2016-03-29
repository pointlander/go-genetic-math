package ast

import "github.com/rogeralsing/GoMath/engine"

//VariableNode represents a variable
type VariableNode struct {
	Variable string
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Eval(context *engine.Context) float64 {
	return context.GetVariable(node.Variable)
}

func (node *VariableNode) ToString() string {
	return node.Variable
}
