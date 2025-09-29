package rsql

import (
	"fmt"
	"strings"
)

type Node interface{}

type LogicalNode struct {
	Operator TokenType
	Children []Node
}

type ComparisonNode struct {
	Field    string
	Operator string
	Value    []string
}

func (n *ComparisonNode) String() string {
	return fmt.Sprintf("(%s %s %s)", n.Field, n.Operator, n.Value)
}

func (n *LogicalNode) String() string {
	var op string
	switch n.Operator {
	case TokenAnd:
		op = "AND"
	case TokenOr:
		op = "OR"
	default:
		op = "UNKNOWN"
	}

	children := make([]string, 0, len(n.Children))
	for _, child := range n.Children {
		children = append(children, fmt.Sprintf("%s", child))
	}

	return fmt.Sprintf("(%s %s)", op, strings.Join(children, " "))
}
