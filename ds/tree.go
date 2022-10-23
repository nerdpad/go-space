package main

// Node
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

type BinaryTree interface {
	Insert(data int)
}

func New() BinaryTree {
	return &Tree{}
}

func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = node(data)
		return
	}

	// find the correct place for this node
}

// findParent finds the parent node for the data.
func findParent(node *Node, data int) *Node {

}

func node(data int) *Node {
	return &Node{Data: data}
}

func main() {

}
