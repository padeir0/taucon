package nodekind

type NodeKind int

func (this NodeKind) String() string {
	switch this {
	case Literal:
		return "Lit."
	case Variable:
		return "Var."
	case BinaryOperator:
		return "Bin. Op."
	case UnaryOperator:
		return "Un. Op."
	case Hole:
		return "_"
	}
	panic("invalid node kind")
}

const (
	InvalidNodeKind NodeKind = iota
	Literal
	Variable
	BinaryOperator
	UnaryOperator

	Hole // we use this to generate
)
