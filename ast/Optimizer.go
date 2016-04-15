package ast

import (
	"github.com/rogeralsing/go-genetic-math/engine"
)

func isLiteralZero(node Node) bool {
	switch t := node.(type) {
	default:
		return false
	case *LiteralNode:
		return t.Value == 0
	}
}

func isLiteral(node Node) bool {
	switch node.(type) {
	default:
		return false
	case *LiteralNode:
		return true
	}
}

func (node *BinaryNode) Optimize() Node {
	left := node.Left.Optimize()
	right := node.Right.Optimize()

	if isLiteral(left) && isLiteral(right) {
		constant := node.Operator.Apply(left, right, engine.EmptyContext)
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
