package generator

import (
	"taucon/tree"
	nk "taucon/tree/nodekind"
)

func Generate(numOps, numVars int) []*tree.Node {
	if numVars > numOps+1 {
		return []*tree.Node{}
	}
	trees := BoolTree(numOps)
	precomputed := map[int][][]int{}
	output := []*tree.Node{}
	for _, t := range trees {
		numHoles := countHoles(t)
		perms, ok := precomputed[numHoles]
		if !ok {
			perms = permute(numVars, numHoles)
			precomputed[numHoles] = perms
		}
		for _, vars := range perms {
			newT := t.Copy() // sorry GC, it must be done
			i := 0
			substitute(newT, vars, &i)
			output = append(output, newT)
		}
	}
	return output
}

func substitute(n *tree.Node, vars []int, i *int) {
	if n.Kind == nk.Hole {
		n.Value = vars[*i]
		n.Kind = nk.Variable
		*i += 1
		return
	}
	if n.Left != nil {
		substitute(n.Left, vars, i)
	}
	if n.Right != nil {
		substitute(n.Right, vars, i)
	}
}

var operators = []int{tree.AND, tree.OR, tree.COND, tree.BICOND}

// Boolean expression tree without leaf information
// 0 -> 1
// 1 -> 5
// 2 -> 45
// 3 -> 505
func BoolTree(numNodes int) []*tree.Node {
	if numNodes == 0 {
		return []*tree.Node{
			{
				Kind:  nk.Hole,
				Value: -1,
			},
		}
	}
	output := []*tree.Node{}
	leftOver := numNodes - 1
	/*
			For a binary node, we combine all trees on the left that are made with N nodes, with all trees in the right that are made with 0 nodes,
			then all trees in the left with N-1 nodes with all trees on the right with 1 Node, until the right = N
			N   0
		    N-1 1
		    N-2 2
		    ... ...
		    1   N-1
		    0   N
	*/
	for i := 0; i <= leftOver; i++ {
		leftNodes := BoolTree(leftOver - i)
		rightNodes := BoolTree(i)
		// this two nested loops combine trees on the left with trees on the right
		for _, leftTree := range leftNodes {
			for _, rightTree := range rightNodes {
				// for each tree shape, we must generate one for all 4 possible operators
				for _, op := range operators {
					newNode := &tree.Node{
						Kind:  nk.BinaryOperator,
						Value: op,
						Left:  leftTree,
						Right: rightTree,
					}
					output = append(output, newNode)
				}
			}
		}
	}
	// for unary trees, there's only one possibility: a single tree with the remaining N nodes
	unaryNodes := BoolTree(leftOver)
	for _, unTree := range unaryNodes {
		newNode := &tree.Node{
			Kind:  nk.UnaryOperator,
			Value: tree.NOT,
			Left:  unTree,
		}
		output = append(output, newNode)
	}
	return output
}

func countHoles(n *tree.Node) int {
	if n.Kind == nk.Hole {
		return 1
	}
	output := 0
	if n.Left != nil {
		output += countHoles(n.Left)
	}
	if n.Right != nil {
		output += countHoles(n.Right)
	}
	return output
}

// generates permutations for variables
func permute(numvars, numslots int) [][]int {
	if numvars == 0 {
		return nil
	}
	output := [][]int{}
	max := pow(numvars, numslots)
	flipfactor := make([]int, numslots)
	for i := 0; i < numslots; i++ {
		flipfactor[i] = max / (pow(numvars, i+1))
	}
	varset := genNums(numvars)
	slots := genNums(numslots)
	for i := 0; fill2(i, max, varset, slots, flipfactor); i++ {
		if containsAll(varset, slots) {
			line := make([]int, len(slots))
			copy(line, slots)
			output = append(output, line)
		}
	}
	return output
}

func genNums(len int) []int {
	out := make([]int, len)
	for i := 0; i < len; i++ {
		out[i] = i
	}
	return out
}

func fill2(index, max int, varset, slots []int, flipfactors []int) bool {
	setlength := len(varset)
	if index >= max {
		return false
	}
	for i := range slots {
		flipfactor := flipfactors[i]
		slots[i] = varset[(index/flipfactor)%setlength]
	}
	return true
}

func containsAll(varset, perm []int) bool {
	m := map[int]struct{}{}
	for _, r := range perm {
		m[r] = struct{}{}
	}
	for _, r := range varset {
		if _, ok := m[r]; !ok {
			return false
		}
	}
	return true
}

func pow(a, b int) int {
	out := 1
	for i := 0; i < b; i++ {
		out *= a
	}
	return out
}
