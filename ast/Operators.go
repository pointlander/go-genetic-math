package ast

import "github.com/rogeralsing/GoMath/engine"
import "math"

//TODO: maybe I should make this mess polymorphic

type BinaryOp int

const (
	OpAdd BinaryOp = iota
	OpSub
	OpDiv
	OpMul
    OpMod
)

var operators = []BinaryOp{
	OpAdd,
	OpSub,
	OpDiv,
    OpMul,	
    OpMod,
}

var operatorSymbols = [...]string{
	"+",
	"-",
	"/",
	"*",
    "%",
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
    func(left Node, right Node, context *engine.Context) float64 {
		return math.Mod(left.Eval(context), right.Eval(context))
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
