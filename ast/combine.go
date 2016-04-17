package ast

func (node *BinaryNode) Combine(other Node) Node {
	if hit(4) {
		return Binary(node.left.Combine(other),node.right,node.operator)
	}
	if hit(4) {
		return Binary(node.left,node.right.Combine(other),node.operator)
	}
	if hit(4) {
		return other
	}
	return node
}

func (node *LiteralNode) Combine(other Node) Node {
	if hit(2) {
		return other
	}
	return node
}

func (node *VariableNode) Combine(other Node) Node {
	if hit(2) {
		return other
	}
	return node
}
