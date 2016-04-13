package ast

import "fmt"
import "strconv"

func (node *BinaryNode) String() string {
    var op string
	switch {
	case OpAdd == node.Operator:
		op = "+"
	case OpSub == node.Operator:
		op = "-"
	case OpMul == node.Operator:
		op = "*"
	case OpDiv == node.Operator:
		op = "/"
	}
    
    return fmt.Sprintf("(%v%v%v)", node.Left, op ,node.Right)
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *VariableNode) String() string {
	return node.Variable
}
