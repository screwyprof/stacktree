package stacktree

import (
	"io"
	"sort"
	"strings"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func PrintStackTrace(input string, w io.Writer) {
	lines := parseLines(input)
	sort.Sort(byLength(lines))
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
		stackTree = stackTree.FindByNameBFS(stackTree, fns[0])
		if stackTree == nil {
			stackTree = New(fns[0], invocations[fns[0]])
		} else {
			stackTree.Invocations = invocations[fns[0]]
		}

		addChild(stackTree, fns[1:], invocations)
	}

	return stackTree
}

func addChild(parent *Node, fns []string, invocations map[string]int) {
	if len(fns) < 1 {
		return
	}
	var child *Node
	child = parent.FindByNameBFS(parent, fns[0])
	if child == nil {
		child = parent.AddChild(fns[0], invocations[fns[0]])
	}
	addChild(child, fns[1:], invocations)
}
