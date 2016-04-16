package domain

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

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

type byFitnessAndWeight []NodeFitness

func (a byFitnessAndWeight) Len() int      { return len(a) }
func (a byFitnessAndWeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byFitnessAndWeight) Less(i, j int) bool {
	if a[i].Fitness < a[j].Fitness {
		return true
	}
	if a[i].Fitness == a[j].Fitness {
		return a[i].Node.Weight() < a[j].Node.Weight()
	}
	return false
}

func calculateFitness(nodes []ast.Node, cases CasesValue) []NodeFitness {
	var fitnessNodes = make([]NodeFitness, len(nodes))
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		fitnessNodes[i].Node = node
		fitnessNodes[i].Fitness = cases.Fitness(node)
	}
	sort.Sort(byFitnessAndWeight(fitnessNodes))
	return fitnessNodes
}

func (cases CasesValue) Solve() ast.Node {

	start := time.Now()
	log.Println(cases)

	populationSize := 20
	generaton := 0
	optimizations := 1000
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

		//create 10 children by genetic crossover
		for i := 0; i < 5; i++ {
			mother := population[rand.Intn(len(population))]
			father := population[rand.Intn(len(population))]
			child := mother.Combine(father)
			population = append(population, child)
		}

		//sort all organisms by fitness
		sorted := calculateFitness(population, cases)
		best := sorted[0]

		if generaton%1000 == 0 {
			log.Printf("Generation %v \t %v  %v", generaton, best.Fitness, best.Node)
		}

		//if we got a better fitness now, print it
		if best.Fitness < bestFitness {
			bestFitness = best.Fitness
		}

		if best.Fitness == 0 {
			optimizations--
		}
		//did we find a solution? if so return it
		if best.Fitness == 0 && optimizations == 0 {
			solution := best.Node.Reduce()
			log.Printf("Solved %v", solution)
			log.Printf("Generations %v", generaton)
			elapsed := time.Since(start)
			log.Printf("Time to find solution %s", elapsed)
			return solution
		}

		population = make([]ast.Node, populationSize)
		for i := 0; i < populationSize; i++ {
			population[i] = sorted[i].Node
		}
		generaton++
	}
}
