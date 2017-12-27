package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Plus) Name() string {
	return "Plus"
}

type Plus struct {
	AssignOp
}

func NewPlus(variable node.Node, expression node.Node) node.Node {
	return Plus{
		AssignOp{
			"AssignPlus",
			variable,
			expression,
		},
	}
}
