package domain

import (
	"github.com/rogeralsing/go-genetic-math/ast"
)

type NodeFitness struct {
	Node    ast.Node
	Fitness float64
}

type byFitnessAndWeight []NodeFitness

func (a byFitnessAndWeight) Len() int {
	return len(a)
}
func (a byFitnessAndWeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byFitnessAndWeight) Less(i, j int) bool {
	if a[i].Fitness < a[j].Fitness {
		return true
	}
	if a[i].Fitness == a[j].Fitness {
		return a[i].Node.Weight() < a[j].Node.Weight()
	}
	return false
}
