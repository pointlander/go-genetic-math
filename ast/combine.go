package ast

func (node *BinaryNode) Combine(other Node) Node {
	if hit(3) {
		return Binary(node.left.Combine(other), node.right, node.operator)
	}
	if hit(3) {
		return Binary(node.left, node.right.Combine(other), node.operator)
	}

	return other.Extract()
}

func (node *LiteralNode) Combine(other Node) Node {
	if hit(2) {
		return other.Extract()
	}
	return node
}

func (node *VariableNode) Combine(other Node) Node {
	if hit(2) {
		return other.Extract()
	}
	return node
}

func (node *BinaryNode) Extract() Node {
	if hit(3) {
		return node.left.Extract()
	}
	if hit(3) {
		return node.right.Extract()
	}

	return node
}

func (node *LiteralNode) Extract() Node {
	return node
}

func (node *VariableNode) Extract() Node {
	return node
}
