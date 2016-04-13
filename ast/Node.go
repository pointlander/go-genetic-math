package ast

import "github.com/rogeralsing/GoMath/engine"

type Node interface {
	Eval(context *engine.Context) float64
	String() string
	Mutate() Node
	IsConstant() bool
	Optimize() Node
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

var operatorSymbols = [...]string{
	"+",
	"-",
	"/",
	"*",
}

var operatorLogic = [...]func(Node, Node, *engine.Context) float64{
	func(left Node, right Node, context *engine.Context) float64 {
		return left.Eval(context) + right.Eval(context)
	},
	func(left Node, right Node, context *engine.Context) float64 {
		return left.Eval(context) - right.Eval(context)
	},
	func(left Node, right Node, context *engine.Context) float64 {
		return left.Eval(context) / right.Eval(context)
	},
	func(left Node, right Node, context *engine.Context) float64 {
		return left.Eval(context) * right.Eval(context)
	},
}

func (op BinaryOp) Apply(left Node, right Node, context *engine.Context) float64 {
	logic := operatorLogic[op]
	res := logic(left, right, context)
	return res
}

func (op BinaryOp) String() string {
	return operatorSymbols[op]
}

type BinaryNode struct {
	Left     Node
	Right    Node
	Operator BinaryOp
}

//VariableNode represents a variable
type VariableNode struct {
	Variable string
}
