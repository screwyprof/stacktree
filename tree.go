package stacktree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Name        string
	Parent      *Node
	Children    []*Node
	Invocations int
}

func New(name string, invocations int) *Node {
	return &Node{Name: name, Invocations: invocations}
}

func (n *Node) Print(w io.Writer) {
	n.traverse(w, n, 0)
}

func (n *Node) traverse(w io.Writer, node *Node, depth int) {
	if node.IsRoot() {
		fmt.Fprintf(w, "%d %s\n", n.Invocations, n.Name)
	}

	for _, child := range node.Children {
		fmt.Fprintf(w, "\t%s%d %s\n", strings.Repeat("\t", depth), child.Invocations, child.Name)
		node.traverse(w, child, depth+1)
	}
}

func (n *Node) AddChild(name string, invocations int) {
	child := New(name, invocations)
	child.setParent(n)
	n.Children = append(n.Children, child)
}

func (n *Node) AddChildNode(child *Node) {
	child.setParent(n)
	n.Children = append(n.Children, child)
}

func (n *Node) setParent(parent *Node) {
	n.Parent = parent
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}
