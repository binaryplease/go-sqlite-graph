package sqlitegraph
type Node struct {
	Id       int
	Children []*Node
	Text     string
}

func NewNode(id int) *Node {
	n := new(Node)
	n.Id = id
	return n
}

// AddChild adds a child to a given node
func (n *Node) AddChild(child *Node) {
}

