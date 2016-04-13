package ast

//Add adds the evaluated result of two nodes
func Add(left Node, right Node) Node {
	return &BinaryNode {Left: left, Right: right, Operator: OpAdd}
}

func Sub(left Node, right Node) Node {
	return &BinaryNode {Left: left, Right: right, Operator: OpSub}
}

func Div(left Node, right Node) Node {
	return &BinaryNode {Left: left, Right: right, Operator: OpDiv}
}

func Mul(left Node, right Node) Node {
	return &BinaryNode {Left: left, Right: right, Operator: OpMul}
}

func Literal(value float64) Node {
	return &LiteralNode{Value: value}
}

func Var(name string) Node {
	return &VariableNode{Variable: name}
}
