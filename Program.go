package main

import (
	"fmt"
	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
)

func main() {
    cases := engine.Cases(engine.Case(20, engine.Input("x",10),engine.Input("y",0)))
    fmt.Printf("%+v",cases)
	context := engine.NewContext()
	context.SetVariable("x", 10)
	add := ast.Add(ast.Mul(ast.Literal(10.111), ast.Literal(20)), ast.Literal(20))
	fmt.Printf("%v", add)
}
