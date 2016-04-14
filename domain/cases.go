package domain

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/rogeralsing/go-genetic-math/ast"
	"github.com/rogeralsing/go-genetic-math/engine"
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

func DefineProblem(cases ...CaseValue) CasesValue {
	return CasesValue{Cases: cases}
}

func (cases CasesValue) String() string {
	var buffer bytes.Buffer
	for x, c := range cases.Cases {
		buffer.WriteString(fmt.Sprintf("Case %v\n", x))
		for _, i := range c.Inputs {
			buffer.WriteString(fmt.Sprintf("\tVar %s = %v \n", i.Variable, i.Value))
		}
		buffer.WriteString(fmt.Sprintf("\tExpected %v \n", c.Result))
	}
	return fmt.Sprint(buffer.String())
}

func (cases CasesValue) Fitness(node ast.Node) float64 {
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

	return total
}

type NodeFitness struct {
	Node    ast.Node
	Fitness float64
}

type byFitness []NodeFitness

func (a byFitness) Len() int           { return len(a) }
func (a byFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool { return a[i].Fitness < a[j].Fitness }

func calculateFitness(nodes []ast.Node, cases CasesValue) []NodeFitness {
	var fitnessNodes = make([]NodeFitness, len(nodes))
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		fitnessNodes[i].Node = node
		fitnessNodes[i].Fitness = cases.Fitness(node)
	}
	sort.Sort(byFitness(fitnessNodes))
	return fitnessNodes
}

func (cases CasesValue) Solve() ast.Node {
	
	fmt.Println(cases)

	populationSize := 5
	var population = make([]ast.Node, populationSize)

	//initialize with dummy data
	for i := 0; i < populationSize; i++ {
		population[i] = ast.Literal(1).Mutate()
	}

	bestFitness := math.MaxFloat64
	for {
		//create a child per parent
		for i := 0; i < populationSize; i++ {
			child := population[i].Mutate()
			population = append(population, child)
		}

		//sort all organisms by fitness
		sorted := calculateFitness(population, cases)
		best := sorted[0]

		//if we got a better fitness now, print it
		if best.Fitness < bestFitness {
			bestFitness = best.Fitness
			fmt.Printf("%v\t%v", best.Fitness, best.Node)
			fmt.Println()
		}

		//did we find a solution? if so return it
		if best.Fitness == 0 {
			fmt.Println()
			fmt.Println("Solved!")
			solution := best.Node.Optimize()
			fmt.Printf("%v",solution)
			return solution
		}

		population = make([]ast.Node, populationSize)
		for i := 0; i < populationSize; i++ {
			population[i] = sorted[i].Node
		}
	}
}
