package ast

import "github.com/rogeralsing/GoMath/engine"
import "fmt"

type MulNode struct {
	Left  Node
	Right Node
}

func (node *MulNode) String() string {
	return fmt.Sprintf("(%v*%v)", node.Left, node.Right)
}

func (node *MulNode) Eval(context *engine.Context) float64 {
	return node.Left.Eval(context) * node.Right.Eval(context)
}

func Mul(left Node, right Node) Node {
	return &MulNode{Left: left, Right: right}
}
