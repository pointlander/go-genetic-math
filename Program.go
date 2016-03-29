package main

import (
	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
    "fmt"
)

func main() {
	context := engine.NewContext()
	context.SetVariable("x", 10)
	add := ast.Add(ast.Mul(ast.Literal(10.111), ast.Literal(20)), ast.Literal(20))
	fmt.Printf("%v",add)    
}
