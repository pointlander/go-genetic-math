package ast

func (node *BinaryNode) IsConstant() bool {
	left := node.Left.IsConstant()
    right := node.Right.IsConstant()
    
    return left && right
}

func (node *LiteralNode) IsConstant() bool {
	return true
}

func (node *VariableNode) IsConstant() bool {
	return false
}
