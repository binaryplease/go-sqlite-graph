package sqlitegraph

// Node is the basic data structure of the graph.
type Node struct {
	ID       int
	Text     string
}

//NewNode creates a new Node with a given id
func NewNode(id int) *Node {
	n := new(Node)
	n.ID = id
	return n
}

func (n *Node) Equals(n2 *Node) bool {
	return (n.ID == n2.ID) && (n.Text == n2.Text)
}
