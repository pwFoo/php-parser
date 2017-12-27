package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n GreaterOrEqual) Name() string {
	return "GreaterOrEqual"
}

type GreaterOrEqual struct {
	BinaryOp
}

func NewGreaterOrEqual(variable node.Node, expression node.Node) node.Node {
	return GreaterOrEqual{
		BinaryOp{
			"BinaryGreaterOrEqual",
			variable,
			expression,
		},
	}
}
