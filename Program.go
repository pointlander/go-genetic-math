package main

import (
	"fmt"

	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
)

func main() {
	cases := engine.Cases(engine.Case(10, engine.Input("x", 5), engine.Input("y", 2)),
		engine.Case(20, engine.Input("x", 10), engine.Input("y", 2)),
		engine.Case(30, engine.Input("x", 10), engine.Input("y", 3)))

	fmt.Printf("%+v", cases)

	add := ast.Add(ast.Mul(ast.Literal(10.111), ast.Literal(20)), ast.Literal(20))

	context := engine.NewContext()
	context.SetVariable("x", 10)

	add = add.Mutate()
	fmt.Printf("%v", add)
}
