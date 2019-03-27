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
	lines := parseLines(input)
	//sort.Sort(byLength(lines))
	stack := buildStackFromFuncs(lines)
	stack.Print(w)
}

func parseLines(input string) []string {
	return strings.Split(input, "\n")
}

func parseALLFuncs(lines []string) []string {
	var fns []string
	for _, line := range lines {
		fns = append(fns, parseFuncs(line)...)
	}
	return fns
}

func parseFuncs(line string) []string {
	fns := strings.Split(line, ",")
	for idx := range fns {
		fns[idx] = strings.TrimSpace(fns[idx])
	}
	return fns
}

func countInvocations(fns []string) map[string]int {
	var invocations = make(map[string]int)
	for i := 0; i < len(fns); i++ {
		invocations[fns[i]]++
	}
	return invocations
}

func buildStackFromFuncs(lines []string) *Node {
	fns := parseALLFuncs(lines)

	var stack *Node
	invocations := countInvocations(fns)

	stack = stack.FindByNameDFS(stack, fns[0])
	if stack == nil {
		stack = New(fns[0], invocations[fns[0]])
	} else {
		stack.Invocations = invocations[fns[0]]
	}
	for i := 1; i < len(fns); i++ {
		child := stack.FindByNameDFS(stack, fns[i])
		if child == nil {
			stack.AddChild(fns[i], invocations[fns[i]])
		}
	}
	return stack
}
