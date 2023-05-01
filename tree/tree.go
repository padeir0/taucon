package tree

import (
	"fmt"
	nk "taucon/tree/nodekind"
)

type Node struct {
	Kind  nk.NodeKind
	Value int // has different meaning depending on NodeKind
	// Literal: Value = 1 is true, Value = 0 is false
	// Variable: Value is the variable id
	// Operators: Value is the operator id, see the "Operators" below

	Left  *Node
	Right *Node
}

// Go doesn't have sum types :)
func (this *Node) IsInvalid() bool {
	switch this.Kind {
	case nk.BinaryOperator:
		if this.Left == nil || this.Right == nil {
			return true
		}
		return this.Left.IsInvalid() || this.Right.IsInvalid()
	case nk.UnaryOperator:
		if this.Left == nil {
			return true
		}
		return this.Left.IsInvalid()
	case nk.Literal:
		// must be a leaf
		if this.Left != nil || this.Right != nil {
			return true
		}
		return this.Value != 0 && this.Value != 1
	case nk.Variable:
		// must be a leaf
		if this.Left != nil || this.Right != nil {
			return true
		}
		return this.Value < 0
	}
	return true // invalid node kind
}

// expression is invalid if a variable is missing
func (this *Node) HasValidVars() bool {
	panic("unimplemented")
}

func (this *Node) String() string {
	switch this.Kind {
	case nk.BinaryOperator:
		op := ""
		switch this.Value {
		case OR:
			op = "∨"
		case AND:
			op = "∧"
		case COND:
			op = "->"
		case BICOND:
			op = "<->"
		}
		return "(" + this.Left.String() + " " + op + " " + this.Right.String() + ")"
	case nk.UnaryOperator:
		if this.Value == NOT {
			return "~" + this.Left.String()
		}
	case nk.Literal:
		if this.Value == 1 {
			return "T"
		} else {
			return "F"
		}
	case nk.Variable:
		return string(rune(this.Value) + 'a')
	case nk.Hole:
		return "_"
	}
	panic("invalid node")
}

func (this *Node) Copy() *Node {
	if this == nil {
		return nil
	}
	left := this.Left.Copy()
	right := this.Right.Copy()
	n := *this
	n.Left = left
	n.Right = right
	return &n
}

func (this *Node) Tree() string {
	return prettyprint(this, 0)
}

func (this *Node) CountOperators() int {
	if this == nil {
		return 0
	}

	total := 0
	if this.Kind == nk.BinaryOperator {
		total = 1
	}
	return total + this.Left.CountOperators() + this.Right.CountOperators()
}

func (this *Node) Hash() int {
	switch this.Kind {
	case nk.BinaryOperator:
		return this.Left.Hash() ^ this.Right.Hash()*(this.Value*31)
	case nk.UnaryOperator:
		return this.Left.Hash() * (this.Value * 31)
	case nk.Variable, nk.Literal:
		return this.Value
	}
	panic("unreachable")
}

func (this *Node) Equal(other *Node) bool {
	if this == nil || other == nil {
		return this == other // only valid if both are nils
	}
	if this.Kind == other.Kind && this.Value == other.Value {
		return this.Left.Equal(other.Left) && this.Right.Equal(other.Right)
	}
	return false
}

func prettyprint(n *Node, i int) string {
	if n == nil {
		return "nil"
	}
	output := fmt.Sprintf("(%v, %v)", n.Kind, n.Value)
	{
		kid := n.Left
		output += indent(i) + prettyprint(kid, i+1)
	}
	{
		kid := n.Right
		output += indent(i) + prettyprint(kid, i+1)
	}

	return output
}

func indent(n int) string {
	output := "\n"
	for i := -1; i < n-1; i++ {
		output += "    "
	}
	output += "└─>"
	return output
}

// Operators
const (
	InvalidOperator int = iota
	AND
	OR
	NOT
	COND
	BICOND
)

func EqualShape(a, b *Node) bool {
	if a != nil && b != nil {
		return EqualShape(a.Left, b.Left) && EqualShape(a.Right, b.Right)
	} else if a == nil && b == nil { // both are nil
		return true
	}
	return false
}
