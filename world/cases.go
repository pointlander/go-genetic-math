package world

import (
	"bytes"
	"fmt"
    "math"
	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
)

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

func (cases CasesValue) Eval(node ast.Node) float64 {
	context := engine.NewContext()
	total := 0.0
	for _, c := range cases.Cases {
		for _, i := range c.Inputs {
			context.SetVariable(i.Variable, i.Value)
		}
		res := node.Eval(context)
		diff := math.Abs(c.Result - res)
		total += diff
	}

	return total // + float64(len(node.String())) / 1000
}
