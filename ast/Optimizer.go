package ast

import (
	"github.com/rogeralsing/go-genetic-math/engine"
)

func isConstantZero(node Node) bool {
	if !node.IsLiteral() {
		return false
	}

	value := node.Eval(engine.NewContext())
	return value == 0
}

func (node *BinaryNode) IsLiteral() bool {
	return false
}

func (node *VariableNode) IsLiteral() bool {
	return false
}

func (node *LiteralNode) IsLiteral() bool {
	return true
}

func (node *BinaryNode) Optimize() Node {
	left := node.Left.Optimize()
	right := node.Right.Optimize()

	if left.IsLiteral() && right.IsLiteral() {
		context := engine.NewContext()
		constant := node.Eval(context)
		return Literal(constant)
	}

	return node.Operator.Optimize(left, right)
}

func (node *LiteralNode) Optimize() Node {
	return node
}

func (node *VariableNode) Optimize() Node {
	return node
}
