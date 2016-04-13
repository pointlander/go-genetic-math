package ast

import (
	"github.com/rogeralsing/GoMath/engine"
)

func (node *BinaryNode) IsConstant() bool {
	left := node.Left.IsConstant()
	right := node.Right.IsConstant()

	return left && right
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
		if left.IsConstant() {
			context := engine.NewContext()
			leftValue := left.Eval(context)

			if leftValue == 0 {
				return right
			}
		}

		if right.IsConstant() {
			context := engine.NewContext()
			rightValue := right.Eval(context)

			if rightValue == 0 {
				return left
			}
		}
	}
    
    //return literal 0 for any multiplication with or by 0
    if node.Operator == OpMul {
        if left.IsConstant() {
			context := engine.NewContext()
			leftValue := left.Eval(context)

			if leftValue == 0 {
				return Literal(0)
			}
		}

		if right.IsConstant() {
			context := engine.NewContext()
			rightValue := right.Eval(context)

			if rightValue == 0 {
				return Literal(0)
			}
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
