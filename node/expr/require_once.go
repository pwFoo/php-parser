package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n RequireOnce) Name() string {
	return "RequireOnce"
}

type RequireOnce struct {
	name string
	expr node.Node
}

func NewRequireOnce(expression node.Node) node.Node {
	return RequireOnce{
		"RequireOnce",
		expression,
	}
}

func (n RequireOnce) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
