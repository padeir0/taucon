package maptree

import (
	"taucon/eval"
	"taucon/generator"
	"taucon/tree"
)

func NewMapTree(numVars int) MapTree {
	trees := []*tree.Node{}
	for i := 0; i <= 3; i++ {
		ts := generator.Generate(i, numVars)
		trees = append(trees, ts...)
	}
	eq := eval.FindEquivalent(trees)
	trees = nil
	return MapTree(eq)
}

type MapTree map[string][]*tree.Node

func (this MapTree) Find(t *tree.Node) ([]*tree.Node, bool) {
	rs := eval.ResultTableOf(t)
	v, ok := this[rs]
	if !ok {
		return nil, false
	}
	return v, true
}
