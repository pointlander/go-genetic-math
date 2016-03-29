package ast

import "github.com/rogeralsing/GoMath/engine"
import "fmt"

type SubNode struct {
	Left  Node
	Right Node
}

func (node *SubNode) String() string {
	return fmt.Sprintf("(%v-%v)", node.Left, node.Right)
}

func Sub(left Node, right Node) Node {
	return &SubNode{Left: left, Right: right}
}

func (node *SubNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) - node.Right.Eval(context)
}
