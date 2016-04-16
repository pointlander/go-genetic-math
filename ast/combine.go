package ast

func (node *BinaryNode) Combine(other Node) Node {
	if hit(4) {
		return node.left.Combine(other)
	}
	if hit(4) {
		return node.right.Combine(other)
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
