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

func (n *Node) AddChild(name string, invocations int) *Node {
	child := New(name, invocations)
	child.setParent(n)
	n.Children = append(n.Children, child)
	return child
}

func (n *Node) AddChildNode(child *Node) {
	child.setParent(n)
	n.Children = append(n.Children, child)
}

func (n *Node) FindByNameBFS(root *Node, name string) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.Name == name {
			return nextUp
		}
		if len(nextUp.Children) > 0 {
			for _, child := range nextUp.Children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

func (n *Node) FindByNameDFS(name string) *Node {
	if n.Name == name {
		return n
	}

	if len(n.Children) > 0 {
		var result *Node
		for i := 0; result == nil && i < len(n.Children); i++ {
			result = n.Children[i].FindByNameDFS(name)
		}
		return result
	}
	return nil
}

func (n *Node) setParent(parent *Node) {
	n.Parent = parent
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}
