package ast

func (this *VariableNode) Mutate() Node {
	if hit(100) {
		copy := &VariableNode{Variable: this.Variable}
		copy.Variable = string(letterRunes[rand.Intn(len(letterRunes))])
	}
	return this
}

func (this *LiteralNode) Mutate() Node {
	//mutate by offset
	if hit(50) {
		copy := &LiteralNode{Value: this.Value}
		copy.Value = this.Value - rand.NormFloat64()*10
		return copy
	} else if hit(50) {
		//hard mutation
		copy := &LiteralNode{Value: this.Value}
		copy.Value = rand.Float64()
		return copy
	}
	return this
}
func (this *AddNode) Mutate() Node {
	return Add(this.Left.Mutate(), this.Right.Mutate())
}
func (this *SubNode) Mutate() Node {
	return Sub(this.Left.Mutate(), this.Right.Mutate())
}
func (this *DivNode) Mutate() Node {
	return Div(this.Left.Mutate(), this.Right.Mutate())
}
func (this *MulNode) Mutate() Node {
	return Mul(this.Left.Mutate(), this.Right.Mutate())
}