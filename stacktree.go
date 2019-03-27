package stacktree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	name        string
	invocations int
	Child       *Node
}

func (n *Node) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "%d %s", n.invocations, n.name)
}

func PrintStackTrace(input string, w io.Writer) {
	var invocations = make(map[string]int)
	var stack *Node

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		invocations[line]++
	}

	stack = &Node{
		name:        lines[0],
		invocations: invocations[lines[0]],
	}

	stack.Print(w)
}
