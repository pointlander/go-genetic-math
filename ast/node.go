package ast

import "github.com/rogeralsing/go-genetic-math/engine"

//Node represents an abstract AST node and the behaviors available on it
type Node interface {
	Eval(*engine.Context) float64
	String() string
	Mutate() Node
	Reduce() Node
	Combine(Node) Node
	Weight() int
	Extract() Node
}

//LiteralNode represents a literal value, e.g. 123.456
type LiteralNode struct {
	value float64
}

//BinaryNode represents a binary operation, e.g. a + b
type BinaryNode struct {
	left     Node
	right    Node
	operator binaryOperator
}

//VariableNode represents a variable, e.g. X
type VariableNode struct {
	variable string
}
