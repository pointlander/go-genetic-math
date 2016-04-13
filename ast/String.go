package ast

import "fmt"
import "strconv"

func (node *BinaryNode) String() string {
	return fmt.Sprintf("(%v%v%v)", node.Left, node.Operator, node.Right)
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *VariableNode) String() string {
	return node.Variable
}
