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
	inputs []InputValue
	result float64
}

func Case(result float64, inputs ...InputValue) CaseValue {
	return CaseValue{
		result: result,
		inputs: inputs,
	}
}

type CasesValue struct {
	cases []CaseValue
}

func DefineProblem(cases ...CaseValue) CasesValue {
	return CasesValue{cases: cases}
}

func (cases CasesValue) String() string {
	var buffer bytes.Buffer
	for x, c := range cases.cases {
		buffer.WriteString(fmt.Sprintf("Case %v\n", x))
		for _, i := range c.inputs {
			buffer.WriteString(fmt.Sprintf("\tVar %s = %v \n", i.Variable, i.Value))
		}
		buffer.WriteString(fmt.Sprintf("\tExpected %v \n", c.result))
	}
	return fmt.Sprint(buffer.String())
}

func (cases CasesValue) Fitness(node ast.Node) float64 {
	context := engine.NewContext()
	total := 0.0
	for _, c := range cases.cases {
		for _, i := range c.inputs {
			context.SetVariable(i.Variable, i.Value)
		}
		res := node.Eval(context)
		diff := math.Abs(c.result - res)
		total += diff
	}

	return total
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
	results := make(chan ast.Node, 1)
	cancel := make(chan struct{}, 1)
	defer close(results)
	defer close(cancel)

	for i := 0; i < 10 ; i++ {
		go solve(cases, results, cancel)
	}

	solution := <-results
	log.Printf("Solved %v", solution)
	elapsed := time.Since(start)
	log.Printf("Time to find solution %s", elapsed)
	return solution
}

func solve(cases CasesValue, results chan<- ast.Node, cancel <-chan struct{}) {
	populationSize := 10
	generaton := 0
	optimizations := 1
	var population = make([]ast.Node, populationSize)

	//initialize with dummy data
	for i := 0; i < populationSize; i++ {
		population[i] = ast.Literal(1).Mutate()
	}

	bestFitness := math.MaxFloat64
	for {

		select {
		case <-cancel:
			log.Println("Quitting")
			return
		default:			
		}

		//create a child per parent
		for i := 0; i < populationSize; i++ {
			child := population[i].Mutate()
			population = append(population, child)
		}

		//create children by genetic crossover
		for i := 0; i < populationSize; i++ {
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
			log.Printf("Generation %v \t %v  %v", generaton, best.Fitness, best.Node)
		}

		if best.Fitness == 0 {
			optimizations--
		}
		//did we find a solution? if so return it
		if best.Fitness == 0 && optimizations == 0 {
			solution := best.Node.Reduce()
			results <- solution
			return
		}

		population = make([]ast.Node, populationSize)
		for i := 0; i < populationSize; i++ {
			population[i] = sorted[i].Node
		}
		generaton++
	}
}
