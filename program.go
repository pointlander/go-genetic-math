package main

import (
	"github.com/rogeralsing/go-genetic-math/domain"
)

func main() {

	// define inputs and expected output
	// problem := domain.DefineProblem(domain.Case(5*2+2, domain.Input("x", 5), domain.Input("y", 2), domain.Input("z", 2)),
	// 	domain.Case(10*2+2, domain.Input("x", 10), domain.Input("y", 2), domain.Input("z", 2)),
	// 	domain.Case(10*3+2, domain.Input("x", 10), domain.Input("y", 3), domain.Input("z", 2)),
	// 	domain.Case(1000*3+5, domain.Input("x", 1000), domain.Input("y", 3), domain.Input("z", 5)),
	// 	domain.Case(333*333+5, domain.Input("x", 333), domain.Input("y", 333), domain.Input("z", 5)))

	//This one is from an old FB spam
	//
	// If we assume that:
	// 2 + 3 = 10
	// 7 + 2 = 63
	// 6 + 5 = 66
	// 8 + 4 = 96

	// How much is?
	// 9 + 7 = ????

	  problem := domain.DefineProblem(domain.Case(10,domain.Input("x",2),domain.Input("y",3)),
	       domain.Case(63,domain.Input("x",7),domain.Input("y",2)),
	       domain.Case(66,domain.Input("x",6),domain.Input("y",5)),
	       domain.Case(96,domain.Input("x",8),domain.Input("y",4)))

	// some binary logic

	// problem := domain.DefineProblem(domain.Case(5&2+2, domain.Input("x", 5), domain.Input("y", 2), domain.Input("z", 2)),
	// 	domain.Case(10&2+2, domain.Input("x", 10), domain.Input("y", 2), domain.Input("z", 2)),
	// 	domain.Case(10&3+2, domain.Input("x", 10), domain.Input("y", 3), domain.Input("z", 2)),
	// 	domain.Case(1000&3+5, domain.Input("x", 1000), domain.Input("y", 3), domain.Input("z", 5)),
	// 	domain.Case(333&333+5, domain.Input("x", 333), domain.Input("y", 333), domain.Input("z", 5)))

	//try to find a formula that matches the above description
	problem.Solve()
}
