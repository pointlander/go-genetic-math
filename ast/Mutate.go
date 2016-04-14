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

var operators = []BinaryOp{
	OpAdd,
	OpSub,
	OpMul,
	OpDiv,
}

func randomOperator() BinaryOp {
	return operators[rand.Intn(len(operators))]
}

var nodeCreators = []func() Node{
	randomLiteralNode,
	randomBinaryNode,
	randomVariableNode,
}

func hit(max int32) bool {
	return rand.Int31n(max) == 1
}

func randomLiteralNode() Node {
	return Literal(rand.NormFloat64())
}

func randomBinaryNode() Node {
	return Binary(randomLiteralNode(), randomLiteralNode(), randomOperator())
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

func randomNode() Node {
	creator := nodeCreators[rand.Intn(len(nodeCreators))]
	node := creator()
	return node
}

func randomRemove(node *BinaryNode) Node {

	if hit(2) {
		return node.Left
	}
	return node.Right
}

func mutateAny(node Node) Node {
	if hit(rate3) {
		return randomNode()
	}
	if hit(rate3) {
		return randomSplit(node)
	}
	if hit(rate1) {
		return node.Optimize()
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
		return Literal(node.Value - (rand.Float64()-0.5)*10)
	}
	if hit(rate1) {
		//hard mutation
		return randomLiteralNode()
	}
	if hit(rate1) {
		//hard mutation to integer

		return Literal(float64(int(node.Value)))
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
	left := node.Left.Mutate()
	right := node.Right.Mutate()
	operator := node.Operator.Mutate()

	mutated := Binary(left, right, operator)

	return mutateAny(mutated)
}

func (operator BinaryOp) Mutate() BinaryOp {
	if hit(rate1) {
		return randomOperator()
	}
	return operator
}
