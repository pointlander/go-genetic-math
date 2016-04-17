package ast

func (node *BinaryNode) Combine(other Node) Node {
	if hit(3) {
		return Binary(node.left.Combine(other), node.right, node.operator)
	}
	if hit(3) {
		return Binary(node.left, node.right.Combine(other), node.operator)
	}

	return other.RandomPart()
}

func (node *LiteralNode) Combine(other Node) Node {
	if hit(2) {
		return other.RandomPart()
	}
	return node
}

func (node *VariableNode) Combine(other Node) Node {
	if hit(2) {
		return other.RandomPart()
	}
	return node
}

func (node *BinaryNode) RandomPart() Node {
	if hit(3) {
		return node.left.RandomPart()
	}
	if hit(3) {
		return node.right.RandomPart()
	}

	return node
}

func (node *LiteralNode) RandomPart() Node {
	return node
}

func (node *VariableNode) RandomPart() Node {
	return node
}
