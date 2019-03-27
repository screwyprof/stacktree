package stacktree

import (
	"io"
	"strings"
)

func PrintStackTrace(input string, w io.Writer) {
	lines := parseLines(input)
	stack := buildStackFromLines(lines)
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

func buildStackFromLines(lines []string) *Node {
	var stackTree *Node
	invocations := countInvocations(parseALLFuncs(lines))

	for _, line := range lines {
		fns := parseFuncs(line)

		// create root
		stackTree = stackTree.FindByNameDFS(stackTree, fns[0])
		if stackTree == nil {
			stackTree = New(fns[0], invocations[fns[0]])
		} else {
			stackTree.Invocations = invocations[fns[0]]
		}

		var child *Node
		for i := 1; i < len(fns); i++ {
			child = stackTree.FindByNameBFS(stackTree, fns[i])
			if child == nil {
				//fmt.Println("child not found")
				stackTree.AddChild(fns[i], invocations[fns[i]])
			} else {
				//fmt.Println("child found")
			}
		}
	}

	return stackTree
}
