package stacktree

import (
	"io"
	"strings"
)

func PrintStackTrace(input string, w io.Writer) {
	var invocations = make(map[string]int)
	var stack *Node

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		fns := strings.Split(line, ",")

		fnsLen := len(fns)
		for i := 0; i < fnsLen; i++ {
			fn := fns[i]
			invocations[fn]++
		}

		stack = New(fns[0], invocations[fns[0]])

		for i := 1; i < fnsLen; i++ {
			stack.AddChild(strings.TrimSpace(fns[i]), invocations[fns[i]])
		}
	}

	stack.Print(w)
}
