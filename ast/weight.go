package ast

//Eval evaluates the result of the binary operation
func (node *BinaryNode) Weight() int {
	return 1 + node.left.Weight() + node.right.Weight()
}

//Eval evaluates the value of the literal node
func (node *LiteralNode) Weight() int {
	return 1
}

//Eval evaluates the value of the variable using a engine.Context
func (node *VariableNode) Weight() int {
	return 1
}
