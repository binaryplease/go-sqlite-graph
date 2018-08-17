package sqlitegraph

// Node is the basic data structure of the graph.
type Node struct {
	ID       int
	Children []*Node
	Text     string
}

//NewNode creates a new Node with a given id
func NewNode(id int) *Node {
	n := new(Node)
	n.ID = id
	return n
}
