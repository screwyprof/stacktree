package stacktree

import (
	"io"
	"sort"
	"strings"
)

func PrintStackTrace(input string, w io.Writer) {
	stack := buildStackFromInput(input)
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

func buildStackFromInput(input string) *Node {
	lines := parseLines(input)
	sort.Sort(byLength(lines))

	var stackTree *Node
	invocations := countInvocations(parseALLFuncs(lines))

	for _, line := range lines {
		fns := parseFuncs(line)

		stackTree = buildTreeRoot(stackTree, fns, invocations)
		addChild(stackTree, fns[1:], invocations)
	}

	return stackTree
}

func buildTreeRoot(stackTree *Node, fns []string, invocations map[string]int) *Node {
	stackTree = stackTree.FindByNameDFS(stackTree, fns[0])
	if stackTree == nil {
		return New(fns[0], invocations[fns[0]])
	}
	stackTree.Invocations = invocations[fns[0]]
	return stackTree
}

func addChild(parent *Node, fns []string, invocations map[string]int) {
	if len(fns) < 1 {
		return
	}
	var child *Node
	child = parent.FindByNameDFS(parent, fns[0])
	if child == nil {
		child = parent.AddChild(fns[0], invocations[fns[0]])
	}
	addChild(child, fns[1:], invocations)
}
