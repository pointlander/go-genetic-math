package ast

import "github.com/rogeralsing/GoMath/engine"

type Node interface {
	Eval(context *engine.Context) float64
	String() string
	Mutate() Node
}

type LiteralNode struct {
	Value float64
}

type BinaryNode struct {
    Left  Node
	Right Node
}

type AddNode struct {
	BinaryNode
}

type DivNode struct {
	BinaryNode
}

type MulNode struct {
	BinaryNode
}

type SubNode struct {
	BinaryNode
}

//VariableNode represents a variable
type VariableNode struct {
	Variable string
}
