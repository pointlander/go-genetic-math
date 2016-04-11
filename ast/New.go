package ast

//Add adds the evaluated result of two nodes
func Add(left Node, right Node) Node {
	return &AddNode{Left: left, Right: right}
}

func Div(left Node, right Node) Node {
	return &DivNode{Left: left, Right: right}
}

func Mul(left Node, right Node) Node {
	return &MulNode{Left: left, Right: right}
}

func Literal(value float64) Node {
	return &LiteralNode{Value: value}
}

func Sub(left Node, right Node) Node {
	return &SubNode{Left: left, Right: right}
}

func Var(name string) Node {
	return &VariableNode{Variable: name}
}
