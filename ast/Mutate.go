package ast

import "math/rand"

const (
	rate1 = 200
	rate2 = 400
	rate3 = 600
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
	return Literal(rand.Float64())
}

func randomBinaryNode() Node {
	return &BinaryNode{Left: randomLiteralNode(), Right: randomLiteralNode(), Operator: randomOperator()}
}

func randomVariableNode() Node {
	return Var(randomLetter())
}

func randomSplit(node Node) Node {
	if hit(1) {
		split := &BinaryNode{Left: randomLiteralNode(), Right: node, Operator: randomOperator()}
		return split
	}
	split := &BinaryNode{Left: node, Right: randomLiteralNode(), Operator: randomOperator()}
	return split
}

func randomNode() Node {
	creator := nodeCreators[rand.Intn(len(nodeCreators))]
	node := creator()
	return node
}

func randomRemove(node *BinaryNode) Node {
	if hit(1) {
		return node.Left
	}
	return node.Right
}

func mutateAny(node Node) Node {
	if hit(rate3) {
		return randomNode()
	} else if hit(rate3) {
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
		return Literal(node.Value - rand.NormFloat64()*10)
	}
	if hit(rate1) {
		//hard mutation
		return Literal(rand.Float64())
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
	if hit(rate1) {
		return &BinaryNode{Left: node.Left.Mutate(), Right: node.Right.Mutate(), Operator: node.Operator}
	}
    //mutate operator
    if hit(rate1) {
		return &BinaryNode{Left: node.Left, Right: node.Right, Operator: randomOperator()}
	}

	return mutateAny(node)
}
