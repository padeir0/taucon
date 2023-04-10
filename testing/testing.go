package testing

import (
	"fmt"
	"taucon/eval"
	"taucon/tree"
	nk "taucon/tree/nodekind"
)

// a ∨ (b ∧ c)
func treeOne() *tree.Node {
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
func treeTwo() *tree.Node {
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
func treeThree() *tree.Node {
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

func Test() {
	expr := treeOne()
	check(expr, eval.Inconstant)

	expr = treeTwo()
	check(expr, eval.True)

	expr = treeThree()
	check(expr, eval.False)
}

func check(expr *tree.Node, expected eval.Result) {
	if eval.ConstValue(expr) != expected {
		fmt.Printf("error: %v was not %v\n", expr, expected)
	}
}
