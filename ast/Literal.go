package ast

import (
	"strconv"

	"github.com/rogeralsing/GoMath/engine"
)

type LiteralNode struct {
	Value float64
}

func Literal(value float64) Node {
	return &LiteralNode{Value: value}
}

func (node *LiteralNode) String() string {
	return strconv.FormatFloat(node.Value, 'f', -1, 64)
}

func (node *LiteralNode) Eval(context *engine.Context) float64 {
	return node.Value
}
