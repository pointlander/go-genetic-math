package ast

import "fmt"
import "strconv"

func (node *AddNode) String() string {
	return fmt.Sprintf("(%v+%v)", node.Left, node.Right)
}

func (node *DivNode) String() string {
	return fmt.Sprintf("(%v/%v)", node.Left, node.Right)
}

func (node *MulNode) String() string {
	return fmt.Sprintf("(%v*%v)", node.Left, node.Right)
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *SubNode) String() string {
	return fmt.Sprintf("(%v-%v)", node.Left, node.Right)
}

func (node *VariableNode) String() string {
	return node.Variable
}
