package ast

import "math/rand"

const (
	rate1 = 50
	rate2 = 100
	rate3 = 250
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
		return Var(string(letterRunes[rand.Intn(len(letterRunes))]))
	}
	return mutateAny(node)
}

//Mutate the given node
func (node *LiteralNode) Mutate() Node {
	//mutate by offset
	if hit(rate1) {
		copy := &LiteralNode{Value: node.Value}
		copy.Value = node.Value - rand.NormFloat64()*10
		return copy
	} 
    if hit(rate1) {
		//hard mutation
		copy := &LiteralNode{Value: node.Value}
		copy.Value = rand.Float64()
		return copy
	}
	return mutateAny(node)
}

//Mutate the given node
func (node *BinaryNode) Mutate() Node {
	if hit(rate1) {
		return Add(node.Left.Mutate(), node.Right.Mutate())
	}
    if hit(rate1) {
        return randomRemove(node)
    }
	return mutateAny(node)
}
