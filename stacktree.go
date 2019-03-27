package stacktree

import (
	"io"
	"strings"
)

//type byLength []string
//
//func (s byLength) Len() int {
//	return len(s)
//}
//func (s byLength) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//func (s byLength) Less(i, j int) bool {
//	return len(s[i]) < len(s[j])
//}

func PrintStackTrace(input string, w io.Writer) {
	var invocations = make(map[string]int)
	var stack *Node

	lines := strings.Split(input, "\n")
	//sort.Sort(byLength(lines))

	for _, line := range lines {
		fns := strings.Split(line, ",")

		fnsLen := len(fns)
		for i := 0; i < fnsLen; i++ {
			fn := fns[i]
			invocations[fn]++
		}

		stack = stack.FindByNameDFS(stack, fns[0])
		if stack == nil {
			stack = New(fns[0], invocations[fns[0]])
		} else {
			stack.Invocations = invocations[fns[0]]
		}

		for i := 1; i < fnsLen; i++ {
			stack.AddChild(strings.TrimSpace(fns[i]), invocations[fns[i]])
		}

	}

	stack.Print(w)
}
