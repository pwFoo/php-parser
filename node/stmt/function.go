package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Function) Name() string {
	return "Function"
}

type Function struct {
	name        string
	token       token.Token
	isReturnRef bool
	params      []node.Node
	returnType  node.Node
	stmts       []node.Node
}

func NewFunction(token token.Token, isReturnRef bool, params []node.Node, returnType node.Node, stmts []node.Node) node.Node {
	return Function{
		"Function",
		token,
		isReturnRef,
		params,
		returnType,
		stmts,
	}
}

func (n Function) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	fmt.Fprintf(out, "\n%vreturn ref: %t", indent+"  ", n.isReturnRef)

	if n.params != nil {
		fmt.Fprintf(out, "\n%vparams:", indent+"  ")
		for _, nn := range n.params {
			nn.Print(out, indent+"    ")
		}
	}

	if n.returnType != nil {
		fmt.Fprintf(out, "\n%vreturn type:", indent+"  ")
		n.returnType.Print(out, indent+"    ")
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
