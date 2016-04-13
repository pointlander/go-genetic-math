package ast

import "github.com/rogeralsing/GoMath/engine"

type Node interface {
	Eval(context *engine.Context) float64
	String() string
	Mutate() Node
    IsConstant() bool
}

type LiteralNode struct {
	Value float64
}

type BinaryOp int

const (
	OpAdd BinaryOp = iota
	OpSub
	OpDiv
	OpMul
)

type BinaryNode struct {
	Left     Node
	Right    Node
	Operator BinaryOp
}

//VariableNode represents a variable
type VariableNode struct {
	Variable string
}
