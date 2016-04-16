package ast

import (
	"github.com/rogeralsing/go-genetic-math/engine"
)

func isLiteralZero(node Node) bool {
	switch t := node.(type) {
	default:
		return false
	case *LiteralNode:
		return t.value == 0
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

func (node *BinaryNode) Reduce() Node {
	left := node.left.Reduce()
	right := node.right.Reduce()

	if isLiteral(left) && isLiteral(right) {
		constant := node.operator.Apply(left, right, engine.EmptyContext)
		return Literal(constant)
	}

	return node.operator.Reduce(left, right)
}

func (node *LiteralNode) Reduce() Node {
	return node
}

func (node *VariableNode) Reduce() Node {
	return node
}
