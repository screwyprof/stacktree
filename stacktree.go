package stacktree

import (
	"io"
	"sort"
	"strings"
)

func PrintStackTrace(input string, w io.Writer) {
	stackTree := buildStackTree(input)
	stackTree.Print(w)
}

func buildStackTree(input string) *Node {
	lines := parseLines(input)
	sort.Sort(byLength(lines))

	stackTree := New("root", 0)
	for idx := range lines {
		fns := parseFuncs(lines[idx])
		invocations := countFnInvocations(fns)

		stackTree = createTree(stackTree, fns, invocations)
		addChild(stackTree, fns[1:], invocations)
	}

	return stackTree
}

func parseLines(input string) []string {
	return strings.Split(input, "\n")
}

func parseFuncs(line string) []string {
	fns := strings.Split(line, ",")
	for idx := range fns {
		fns[idx] = strings.TrimSpace(fns[idx])
	}
	return fns
}

func countFnInvocations(fns []string) map[string]int {
	var invocations = make(map[string]int)
	for i := 0; i < len(fns); i++ {
		fn := fns[i]
		invocations[fn]++
	}
	return invocations
}

func createTree(stackTree *Node, fns []string, invocations map[string]int) *Node {
	stackTree = stackTree.FindByNameDFS(fns[0])
	if stackTree == nil {
		stackTree = New(fns[0], 0)
	}
	stackTree.Invocations += invocations[fns[0]]
	return stackTree
}

func addChild(parent *Node, fns []string, invocations map[string]int) {
	var child *Node
	for i := 0; i < len(fns); i++ {
		child = parent.FindByNameDFS(fns[i])
		if child == nil {
			child = parent.AddChild(fns[i], 0)
		}
		child.Invocations += invocations[fns[i]]
		parent = child
	}
}
