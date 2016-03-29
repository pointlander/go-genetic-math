package ast

import "github.com/rogeralsing/GoMath/engine"
import "fmt"

type DivNode struct {
	Left  Node
	Right Node
}

func (node *DivNode) String() string {
	return fmt.Sprintf("(%v/%v)",node.Left,node.Right)
}

func Div(left Node, right Node) Node {
	return &DivNode{Left: left, Right: right}
}

func (node *DivNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) / node.Right.Eval(context)
}
