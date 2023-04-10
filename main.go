package main

import (
	"fmt"
	"taucon/eval"
	"taucon/generator"
	"taucon/tree"
	nk "taucon/tree/nodekind"
)

func main() {
	printConst()
}

// a ∨ (b ∧ c)
func TreeOne() *tree.Node {
	a := &tree.Node{
		Kind:  nk.Variable,
		Value: 0,
	}
	b := &tree.Node{
		Kind:  nk.Variable,
		Value: 1,
	}
	c := &tree.Node{
		Kind:  nk.Variable,
		Value: 2,
	}
	and := &tree.Node{
		Kind:  nk.BinaryOperator,
		Value: tree.AND,
		Left:  b,
		Right: c,
	}
	or := &tree.Node{
		Kind:  nk.BinaryOperator,
		Value: tree.OR,
		Left:  a,
		Right: and,
	}
	return or
}

// a ∨ ~a
func TreeTwo() *tree.Node {
	a := &tree.Node{
		Kind:  nk.Variable,
		Value: 0,
	}
	not := &tree.Node{
		Kind:  nk.UnaryOperator,
		Value: tree.NOT,
		Left:  a,
	}
	or := &tree.Node{
		Kind:  nk.BinaryOperator,
		Value: tree.OR,
		Left:  a,
		Right: not,
	}
	return or
}

// a ∧ ~a
func TreeThree() *tree.Node {
	a := &tree.Node{
		Kind:  nk.Variable,
		Value: 0,
	}
	not := &tree.Node{
		Kind:  nk.UnaryOperator,
		Value: tree.NOT,
		Left:  a,
	}
	and := &tree.Node{
		Kind:  nk.BinaryOperator,
		Value: tree.AND,
		Left:  a,
		Right: not,
	}
	return and
}

func test() {
	fmt.Println("--------------------")
	expr := TreeOne()
	fmt.Println(expr)
	fmt.Println(eval.TruthTableOf(expr))
	fmt.Println(eval.ConstValue(expr))

	fmt.Println("--------------------")
	expr = TreeTwo()
	fmt.Println(expr)
	fmt.Println(eval.TruthTableOf(expr))
	fmt.Println(eval.ConstValue(expr))

	fmt.Println("--------------------")
	expr = TreeThree()
	fmt.Println(expr)
	fmt.Println(eval.TruthTableOf(expr))
	fmt.Println(eval.ConstValue(expr))
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

type ConstTree struct {
	Expr   *tree.Node
	Result bool
}

func (this ConstTree) String() string {
	res := "F"
	if this.Result {
		res = "T"
	}
	return this.Expr.String() + " = " + res
}

func printConst() {
	taucon := []ConstTree{}
	trees := generator.Generate(3, 2)
	for _, t := range trees {
		res := eval.ConstValue(t)
		switch res {
		case eval.True:
			taucon = append(taucon, ConstTree{
				Expr:   t,
				Result: true,
			})
		case eval.False:
			taucon = append(taucon, ConstTree{
				Expr:   t,
				Result: false,
			})
		}
	}
	for _, ct := range taucon {
		fmt.Println(ct)
	}
}
