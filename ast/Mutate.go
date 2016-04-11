package ast

import "math/rand"

const (
	rate1 = 50
	rate2 = 100
	rate3 = 250
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

var creators = []func() Node{
	randomLiteralNode,
	randomAddNode,
	randomSubNode,
	randomMulNode,
	randomDivNode,
	randomVariableNode,
}

func hit(max int32) bool {
	return rand.Int31n(max) == 1
}

func randomLiteralNode() Node {
	return Literal(rand.Float64())
}

func randomAddNode() Node {
	return Add(randomLiteralNode(), randomLiteralNode())
}

func randomSubNode() Node {
	return Sub(randomLiteralNode(), randomLiteralNode())
}

func randomDivNode() Node {
	return Div(randomLiteralNode(), randomLiteralNode())
}

func randomMulNode() Node {
	return Mul(randomLiteralNode(), randomLiteralNode())
}

func randomVariableNode() Node {
	return Var(string(letterRunes[rand.Intn(len(letterRunes))]))
}

func randomNode() Node {
	creator := creators[rand.Intn(len(creators))]
	node := creator()
	return node
}

//Mutate the given node
func (node *VariableNode) Mutate() Node {
	if hit(rate2) {
		copy := &VariableNode{Variable: node.Variable}
		copy.Variable = string(letterRunes[rand.Intn(len(letterRunes))])
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}

//Mutate the given node
func (node *LiteralNode) Mutate() Node {
	//mutate by offset
	if hit(rate1) {
		copy := &LiteralNode{Value: node.Value}
		copy.Value = node.Value - rand.NormFloat64()*10
		return copy
	} else if hit(rate1) {
		//hard mutation
		copy := &LiteralNode{Value: node.Value}
		copy.Value = rand.Float64()
		return copy
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}

//Mutate the given node
func (node *AddNode) Mutate() Node {
	if hit(rate1) {
		return Add(node.Left.Mutate(), node.Right.Mutate())
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}

//Mutate the given node
func (node *SubNode) Mutate() Node {
	if hit(rate1) {
		return Sub(node.Left.Mutate(), node.Right.Mutate())
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}

//Mutate the given node
func (node *DivNode) Mutate() Node {
	if hit(rate1) {
		return Div(node.Left.Mutate(), node.Right.Mutate())
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}

//Mutate the given node
func (node *MulNode) Mutate() Node {
	if hit(rate1) {
		return Mul(node.Left.Mutate(), node.Right.Mutate())
	} else if hit(rate3) {
		return randomNode()
	}
	return node
}
