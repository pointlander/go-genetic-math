package engine

import (
	"bytes"
	"fmt"
)

type Context struct {
	variables map[string]float64
}

func NewContext() *Context {
	return &Context{
		variables: map[string]float64{},
	}
}

func (node *Context) GetVariable(variable string) float64 {
	return node.variables[variable]
}

func (node *Context) SetVariable(variable string, value float64) {
	node.variables[variable] = value
}

type InputValue struct {
	Variable string
	Value    float64
}

func Input(variable string, value float64) InputValue {
	return InputValue{Variable: variable, Value: value}
}

type CaseValue struct {
	Inputs []InputValue
	Result float64
}

func Case(result float64, inputs ...InputValue) CaseValue {
	return CaseValue{
		Result: result,
		Inputs: inputs,
	}
}

type CasesValue struct {
	Cases []CaseValue
}

func Cases(cases ...CaseValue) CasesValue {
	return CasesValue{Cases: cases}
}

func (cases CasesValue) String() string {
	var buffer bytes.Buffer
	for x, c := range cases.Cases {
		buffer.WriteString(fmt.Sprintf("Case %v = %v \n", x, c.Result))
		for _, i := range c.Inputs {
			buffer.WriteString(fmt.Sprintf("\tVar %s = %v \n", i.Variable, i.Value))
		}
	}
	return fmt.Sprint(buffer.String())
}
