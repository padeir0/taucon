package eval

import (
	"taucon/tree"
	nk "taucon/tree/nodekind"
)

type Result byte

func (this Result) String() string {
	switch this {
	case True:
		return "tautology"
	case False:
		return "contradiction"
	case Inconstant:
		return "not constant"
	}
	return "?"
}
func (this Result) Pretty() string {
	switch this {
	case True:
		return "V"
	case False:
		return "F"
	case Inconstant:
		return "z"
	}
	return "?"
}

const (
	Inconstant Result = iota
	True
	False
	unset
)

func ConstValue(n *tree.Node) Result {
	numvars := findVars(n, 0)
	vars := make([]bool, numvars)
	res := unset
	for i := 0; fillVariables(i, &vars); i++ {
		t := eval(n, vars)
		if res == unset {
			if t {
				res = True
			} else {
				res = False
			}
		} else {
			if (res == True && !t) || (res == False && t) {
				return Inconstant
			}
		}
	}
	return res
}

// returns the constant value of an expression, if any
func eval(n *tree.Node, vars []bool) bool {
	switch n.Kind {
	case nk.Literal:
		return n.Value == 1
	case nk.Variable:
		return vars[n.Value]
	case nk.BinaryOperator:
		switch n.Value {
		case tree.AND:
			return eval(n.Left, vars) && eval(n.Right, vars)
		case tree.OR:
			return eval(n.Left, vars) || eval(n.Right, vars)
		case tree.COND:
			if eval(n.Left, vars) {
				return eval(n.Right, vars)
			}
			return true
		case tree.BICOND:
			return eval(n.Left, vars) == eval(n.Right, vars)
		}
		panic("invalid operator")
	case nk.UnaryOperator:
		if n.Value == tree.NOT {
			return !eval(n.Left, vars)
		}
		panic("invalid operator")
	}
	panic("invalid expression")
}

type TruthTable [][]bool

func (this TruthTable) String() string {
	output := ""
	if len(this) > 0 {
		for i := 0; i < len(this[0])-1; i++ {
			output += string(rune(i)+'a') + "  "
		}
		output += "expr \n"
	}
	for _, row := range this {
		for _, b := range row {
			if b {
				output += "V, "
			} else {
				output += "F, "
			}
		}
		output += "\n"
	}
	return output
}

func FindEquivalent(trees []*tree.Node) map[string][]*tree.Node {
	eq := map[string][]*tree.Node{}
	for _, t := range trees {
		res := ResultTableOf(t)
		v, ok := eq[res]
		if !ok {
			eq[res] = []*tree.Node{t}
		} else {
			eq[res] = append(v, t)
		}
	}
	// we keep only the equivalent ones, to save memory, odds are most are kept anyway
	// because of commutativity and associativity
	toRemove := []string{}
	for k, v := range eq {
		if len(v) < 1 {
			toRemove = append(toRemove, k)
		}
	}
	for _, rm := range toRemove {
		delete(eq, rm)
	}
	return eq
}

func ResultTableOf(n *tree.Node) string {
	numvars := findVars(n, 0)
	output := ""
	vars := make([]bool, numvars)
	for i := 0; fillVariables(i, &vars); i++ {
		t := eval(n, vars)
		output += boolToS(t)
	}
	return output
}

func boolToS(b bool) string {
	if b {
		return "V"
	} else {
		return "F"
	}
}

// returns the truth table
func TruthTableOf(n *tree.Node) TruthTable {
	numvars := findVars(n, 0)
	output := make([][]bool, 1<<numvars)
	vars := make([]bool, numvars)
	for i := 0; fillVariables(i, &vars); i++ {
		t := eval(n, vars)
		cp := make([]bool, numvars)
		copy(cp, vars)
		cp = append(cp, t)
		output[i] = cp
	}
	return output
}

func Test(numvars int) TruthTable {
	output := make([][]bool, 1<<numvars)
	vars := make([]bool, numvars)
	for i := 0; fillVariables(i, &vars); i++ {
		cp := make([]bool, numvars)
		copy(cp, vars)
		output[i] = cp
	}
	return output
}

func findVars(n *tree.Node, count int) int {
	if n.Kind == nk.Variable {
		newcount := n.Value + 1
		if newcount > count {
			return newcount
		}
		return count
	}
	newcount := count
	if n.Left != nil {
		newcount = findVars(n.Left, newcount)
	}
	if n.Right != nil {
		newcount = findVars(n.Right, newcount)
	}
	return newcount
}

func fillVariables(index int, vars *[]bool) bool {
	numvars := len(*vars)
	max := 1 << numvars
	if index >= max {
		return false
	}
	for i := range *vars {
		flipfactor := max / (1 << (i + 1))
		if index%flipfactor == 0 {
			(*vars)[i] = !(*vars)[i]
		}
	}
	return true
}
