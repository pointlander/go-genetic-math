package main

import (
	"github.com/rogeralsing/GoMath/domain"
)

func main() {
    
    //define inputs and expected output
	problem := domain.DefineProblem(domain.Case(5*2+2, domain.Input("x", 5), domain.Input("y", 2), domain.Input("z", 2)),
		domain.Case(10*2+2, domain.Input("x", 10), domain.Input("y", 2), domain.Input("z", 2)),
		domain.Case(10*3+2, domain.Input("x", 10), domain.Input("y", 3), domain.Input("z", 2)),
		domain.Case(1000*3+5, domain.Input("x", 1000), domain.Input("y", 3), domain.Input("z", 5)),
		domain.Case(333*333+5, domain.Input("x", 333), domain.Input("y", 333), domain.Input("z", 5)))

    //try to find a formula that matches the above description
	node := problem.Solve()
	
	println(node.String())
}
