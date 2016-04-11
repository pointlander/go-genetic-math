package main

import (
	"fmt"

	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
	"github.com/rogeralsing/GoMath/world"
)

func main() {
	cases := world.Cases(world.Case(10, world.Input("x", 5), world.Input("y", 2)),
		world.Case(20, world.Input("x", 10), world.Input("y", 2)),
		world.Case(30, world.Input("x", 10), world.Input("y", 3)),
        world.Case(3000, world.Input("x", 1000), world.Input("y", 3)))

	fmt.Printf("%+v", cases)

	add := ast.Add(ast.Mul(ast.Literal(10.111), ast.Literal(20)), ast.Literal(20))

	context := engine.NewContext()
	context.SetVariable("x", 10)

	parent := add
	parentFitness := cases.Eval(parent)
	for {
		child := parent.Mutate()
		childFitness := cases.Eval(child)
		if childFitness < parentFitness {
			parent = child
			parentFitness = childFitness

			fmt.Printf("%v\t%v", parent, parentFitness)
            println()
		}
	}
}
