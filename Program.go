package main

import (
	"github.com/rogeralsing/GoMath/world"
)

func main() {
	cases := world.Cases(world.Case(5*2+2, world.Input("x", 5), world.Input("y", 2), world.Input("z",2)),
		world.Case(10*2+2, world.Input("x", 10), world.Input("y", 2), world.Input("z",2)),
		world.Case(10*3+2, world.Input("x", 10), world.Input("y", 3), world.Input("z",2)),
		world.Case(1000*3+5, world.Input("x", 1000), world.Input("y", 3), world.Input("z",5)),
		world.Case(333*333+5, world.Input("x", 333), world.Input("y", 333), world.Input("z",5)))

	node := cases.Solve()
    node = node.Optimize()
	println(node.String())
}
