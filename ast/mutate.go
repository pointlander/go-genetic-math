package ast

import "math/rand"

const (
	rate1 = 50
	rate2 = 100
	rate3 = 200
)

var letterRunes = []rune("xyz")

func randomLetter() string {
	return string(letterRunes[rand.Intn(len(letterRunes))])
}

func randomOperator() BinaryOp {
	return operators[rand.Intn(len(operators))]
}

func hit(max int32) bool {
	return rand.Int31n(max) == 1
}

func randomLiteralNode() Node {
	return Literal(rand.NormFloat64())
}

func randomBinaryNode() Node {
	return Binary(randomNode(), randomNode(), randomOperator())
}

func randomVariableNode() Node {
	return Var(randomLetter())
}

func randomSplit(node Node) Node {
	if hit(2) {
		split := Binary(randomLiteralNode(), node, randomOperator())
		return split
	}
	split := Binary(node, randomLiteralNode(), randomOperator())
	return split
}

//Prio
//VariableNode
//BinaryNode
//LiteralNode
//
//this is because we most likely want a formula based on variables, and not a lot of magic constants
func randomNode() Node {
	if hit(500) {
		return randomLiteralNode()
	}
	if hit(20) {
		return randomBinaryNode()
	}
	return randomVariableNode()
}

func randomRemove(node *BinaryNode) Node {

	if hit(2) {
		return node.left
	}
	return node.right
}

func mutateAny(node Node) Node {
	if hit(rate3) {
		return randomNode()
	}
	if hit(rate3) {
		return randomSplit(node)
	}

	return node
}

//Mutate the given node
func (node *VariableNode) Mutate() Node {
	if hit(rate2) {
		return Var(randomLetter())
	}
	return mutateAny(node)
}

//Mutate the given node
func (node *LiteralNode) Mutate() Node {
	//mutate by offset
	if hit(rate1) {
		return Literal(node.value - (rand.Float64()-0.5)*10)
	}
	if hit(rate1) {
		//hard mutation
		return randomLiteralNode()
	}
	if hit(rate1) {
		//hard mutation to integer
		return Literal(float64(int(node.value)))
	}

	return mutateAny(node)
}

//Mutate the given node
func (node *BinaryNode) Mutate() Node {
	//remove left or right
	if hit(rate1) {
		return randomRemove(node)
	}
	//mutate children
	left := node.left.Mutate()
	right := node.right.Mutate()

	operator := node.operator
	if hit(rate1) {
		operator = randomOperator()
	}

	mutated := Binary(left, right, operator)

	return mutateAny(mutated)
}
