package ast

import "github.com/rogeralsing/GoMath/engine"
import "fmt"

//AddNode represent an addition
type AddNode struct {
	Left  Node
	Right Node
}

//Add adds the evaluated result of two nodes
func Add(left Node, right Node) Node {
	return &AddNode{Left: left, Right: right}
}

func (node *AddNode) String() string {
	return fmt.Sprintf("(%v+%v)",node.Left,node.Right)
}

func (node *AddNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) + node.Right.Eval(context)
}
