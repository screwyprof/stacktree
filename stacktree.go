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

func (n *Node) AddChild(child *Node) {
	n.Child = child
}

func PrintStackTrace(input string, w io.Writer) {
	var invocations = make(map[string]int)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		invocations[line]++
	}

	_, _ = fmt.Fprintf(w, "%d %s", invocations[lines[0]], lines[0])
}
