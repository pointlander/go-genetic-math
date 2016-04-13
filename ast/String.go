package ast

import "fmt"
import "strconv"

func (node *BinaryNode) String() string {
	switch {
	case OpAdd == node.Operator:
		return fmt.Sprintf("(%v+%v)", node.Left, node.Right)
	case OpSub == node.Operator:
		return fmt.Sprintf("(%v-%v)", node.Left, node.Right)
	case OpMul == node.Operator:
		return fmt.Sprintf("(%v*%v)", node.Left, node.Right)
	case OpDiv == node.Operator:
		return fmt.Sprintf("(%v/%v)", node.Left, node.Right)
	default:
		return "unknown"
	}
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *VariableNode) String() string {
	return node.Variable
}
