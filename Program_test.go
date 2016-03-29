package main

import (
	"testing"

	"github.com/rogeralsing/GoMath/ast"
	"github.com/rogeralsing/GoMath/engine"
	"github.com/stretchr/testify/assert"
)

func TestBar(t *testing.T) {
	context := engine.NewContext()
	context.SetVariable("hej", 123)
	add := ast.Add(ast.Literal(2), ast.Literal(3))
	res := add.Eval(context)
	println(res)
	assert.Equal(t, 5.0, res, "should be 5")
}
