package ast

import "strconv"

func (node *BinaryNode) String() string {
	return node.Operator.String(node.Left, node.Right)
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *VariableNode) String() string {
	return node.Variable
}
