package main

import (
	"fmt"
	"taucon/eval"
	"taucon/generator"
	"taucon/maptree"
	"taucon/tree"
	"time"
)

func main() {
	//countConst()
	printEquivalent()
	//PrintConst()
}

func gentest() {
	{
		trees := generator.Generate(2, 1)
		fmt.Println("Amount: ", len(trees))
		for _, t := range trees {
			fmt.Println(t)
		}
	}
	fmt.Println("f(numOps, numVars)")
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			trees := generator.Generate(i, j)
			fmt.Printf("f(%v, %v) = %v\n", i, j, len(trees))
		}
	}
}

func countConst() {
	fmt.Println("f(i, j) where i = number of operators, j = number of variables")
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			F(i, j)
		}
	}
}

func printEquivalent() {
	mp := maptree.NewMapTree(3)
	for _, group := range mp {
		fmt.Println("-----------------------------------")
		for _, t := range group {
			fmt.Println(t)
		}
	}
}

func F(i, j int) {
	tau := []*tree.Node{}
	con := []*tree.Node{}
	inconst := []*tree.Node{}
	start := time.Now()
	trees := generator.Generate(i, j)
	duration := time.Since(start)
	for _, t := range trees {
		res := eval.ConstValue(t)
		switch res {
		case eval.True:
			tau = append(tau, t)
		case eval.False:
			con = append(con, t)
		default:
			inconst = append(inconst, t)
		}
	}
	fmt.Printf("%v\t\tf(%v, %v) = (inconst: %v, const: %v, con: %v, tau: %v, total: %v)\n", duration, i, j, len(inconst), len(tau)+len(con), len(con), len(tau), len(tau)+len(con)+len(inconst))
}

func PrintConst() {
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			printTrees(i, j)
		}
	}
}

func printTrees(i, j int) {
	fmt.Printf("f(%v, %v)-------------------------------\n", i, j)
	trees := generator.Generate(i, j)
	for _, t := range trees {
		res := eval.ConstValue(t)
		switch res {
		case eval.True, eval.False:
			fmt.Println(t, " <=> ", res.Pretty())
		}
	}
}
