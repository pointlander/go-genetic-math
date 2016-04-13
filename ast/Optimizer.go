package ast

import (
	"github.com/rogeralsing/GoMath/engine"
)

func isConstantZero(node Node) bool {
	if !node.IsConstant() {
        return false
    }
    
	value := node.Eval(engine.NewContext())
	return value == 0
}

func (node *BinaryNode) IsConstant() bool {

	left := node.Left
	right := node.Right

	//if this is a mul node and one of the operands are 0, then this is constant
	if node.Operator == OpMul {
		if isConstantZero(left) {
			return true
		}

		if isConstantZero(right) {
			return true
		}
	}

	return node.Left.IsConstant() && node.Right.IsConstant()
}

func (node *LiteralNode) IsConstant() bool {
	return true
}

func (node *VariableNode) IsConstant() bool {
	return false
}

func (node *BinaryNode) Optimize() Node {

	//if the entre node is constant, evaluate and return literal with content
	if node.IsConstant() {
		context := engine.NewContext()
		constant := node.Eval(context)
		return Literal(constant)
	}

	left := node.Left.Optimize()
	right := node.Right.Optimize()

	//remove any + or - of constant 0
	if node.Operator == OpAdd || node.Operator == OpSub {
		if isConstantZero(left) {
			return right
		}

		if isConstantZero(right) {
			return left
		}
	}

	//return literal 0 for any multiplication with or by 0
	if node.Operator == OpMul {
		if isConstantZero(left) {
			return Literal(0)
		}

		if isConstantZero(right) {
			return Literal(0)
		}
	}

	return Binary(left, right, node.Operator)
}

func (node *LiteralNode) Optimize() Node {
	return node
}

func (node *VariableNode) Optimize() Node {
	return node
}
