package ast

import "github.com/rogeralsing/go-genetic-math/engine"
import "math"
import "fmt"

//TODO: maybe I should make this mess polymorphic

type BinaryOp interface {
	Apply(left Node, right Node, context *engine.Context) float64
	String(left Node, right Node) string
	Optimize(left Node, right Node) Node
}

type OpAddValue struct{}
type OpSubValue struct{}
type OpDivValue struct{}
type OpMulValue struct{}
type OpModValue struct{}
type OpOrValue struct{}
type OpAndValue struct{}

var OpAdd = OpAddValue{}
var OpSub = OpSubValue{}
var OpDiv = OpDivValue{}
var OpMul = OpMulValue{}
var OpMod = OpModValue{}
var OpOr = OpOrValue{}
var OpAnd = OpAndValue{}

var operators = []BinaryOp{
	OpAdd,
	OpSub,
	OpMul,
	OpDiv,
	OpMod,
	OpOr,
	OpAnd,
	// OpXor,
}

func (OpAddValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return left.Eval(context) + right.Eval(context)
}
func (OpSubValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return left.Eval(context) - right.Eval(context)
}
func (OpDivValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return left.Eval(context) / right.Eval(context)
}
func (OpMulValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return left.Eval(context) * right.Eval(context)
}
func (OpModValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return math.Mod(left.Eval(context), right.Eval(context))
}
func (OpOrValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return float64(int(left.Eval(context)) | int(right.Eval(context)))
}
func (OpAndValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return float64(int(left.Eval(context)) & int(right.Eval(context)))
}

func (OpAddValue) String(left Node, right Node) string {
	return fmt.Sprintf("(%v+%v)", left, right)
}
func (OpSubValue) String(left Node, right Node) string {
	return fmt.Sprintf("(%v-%v)", left, right)
}
func (OpDivValue) String(left Node, right Node) string {
	return fmt.Sprintf("(%v/%v)", left, right)
}
func (OpMulValue) String(left Node, right Node) string {
	return fmt.Sprintf("(%v*%v)", left, right)
}
func (OpModValue) String(left Node, right Node) string {
	return fmt.Sprintf("mod(%v,%v)", left, right)
}
func (OpOrValue) String(left Node, right Node) string {
	return fmt.Sprintf("((int)%v|(int)%v)", left, right)
}
func (OpAndValue) String(left Node, right Node) string {
	return fmt.Sprintf("((int)%v&(int)%v)", left, right)
}

func (operator OpAddValue) Optimize(left Node, right Node) Node {
	if isConstantZero(left) {
		return right
	}

	if isConstantZero(right) {
		return left
	}

	return Binary(left, right, operator)
}
func (operator OpSubValue) Optimize(left Node, right Node) Node {
	if isConstantZero(left) {
		return right
	}

	if isConstantZero(right) {
		return left
	}

	return Binary(left, right, operator)
}
func (operator OpDivValue) Optimize(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpMulValue) Optimize(left Node, right Node) Node {
	if isConstantZero(left) || isConstantZero(right) {
		return Literal(0)
	}
	return Binary(left, right, operator)
}
func (operator OpModValue) Optimize(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpOrValue) Optimize(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpAndValue) Optimize(left Node, right Node) Node {
	return Binary(left, right, operator)
}
