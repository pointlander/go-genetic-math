package ast

import "strconv"

func (node *BinaryNode) String() string {
	return node.operator.String(node.left, node.right)
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.value, 'f', -1, 64)
}

func (node *VariableNode) String() string {
	return node.variable
}
