package ast

import "github.com/rogeralsing/go-genetic-math/engine"
import "math"
import "fmt"

type binaryOperator interface {
	Apply(left Node, right Node, context *engine.Context) float64
	String(left Node, right Node) string
	Reduce(left Node, right Node) Node
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

var operators = [...]binaryOperator{
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
	leftValue := left.Eval(context)
	rightValue := right.Eval(context)
	if rightValue == 0 {
		return 0
	}
	return leftValue / rightValue
}
func (OpMulValue) Apply(left Node, right Node, context *engine.Context) float64 {
	return left.Eval(context) * right.Eval(context)
}
func (OpModValue) Apply(left Node, right Node, context *engine.Context) float64 {
	leftValue := left.Eval(context)
	rightValue := right.Eval(context)
	if rightValue <= 0 {
		return 0
	}
	return math.Mod(leftValue, rightValue)
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

func (operator OpAddValue) Reduce(left Node, right Node) Node {

	//if left is 0, we can reduce this to only the right node
	if isLiteralZero(left) {
		return right
	}

	//if right is 0, we can reduce this to only the left node
	if isLiteralZero(right) {
		return left
	}

	return Binary(left, right, operator)
}
func (operator OpSubValue) Reduce(left Node, right Node) Node {
	//if left is 0, we can reduce this to only the right node
	if isLiteralZero(left) {
		return right
	}

	//if right is 0, we can reduce this to only the left node
	if isLiteralZero(right) {
		return left
	}

	return Binary(left, right, operator)
}
func (operator OpDivValue) Reduce(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpMulValue) Reduce(left Node, right Node) Node {

	//anything multiplied by 0 is 0, reduce to constant
	if isLiteralZero(left) || isLiteralZero(right) {
		return Literal(0)
	}
	return Binary(left, right, operator)
}
func (operator OpModValue) Reduce(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpOrValue) Reduce(left Node, right Node) Node {
	return Binary(left, right, operator)
}
func (operator OpAndValue) Reduce(left Node, right Node) Node {
	return Binary(left, right, operator)
}
