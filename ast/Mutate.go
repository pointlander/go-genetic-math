package ast

import "math/rand"

const (
	rate1 = 150
	rate2 = 200
	rate3 = 350
)

var letterRunes = []rune("xyz")

var randomBinaryCreator = []func(Node, Node) Node{
	Add,
	Sub,
	Mul,
	Div,
}

var randomNodeCreators = []func() Node{
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

func randomSplit(node Node) Node {
	creator := randomBinaryCreator[rand.Intn(len(randomBinaryCreator))]
	if hit(1) {
		split := creator(node, randomLiteralNode())
		return split
	}
	split := creator(randomLiteralNode(), node)
	return split
}

func randomNode() Node {
	creator := randomNodeCreators[rand.Intn(len(randomNodeCreators))]
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
        variable := string(letterRunes[rand.Intn(len(letterRunes))])
		return Var(variable)
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
	if hit(rate1) {
		return &BinaryNode { Left: node.Left.Mutate(), Right: node.Right.Mutate(), Operator: node.Operator}
	}
    if hit(rate1) {
        return randomRemove(node)
    }
	return mutateAny(node)
}
