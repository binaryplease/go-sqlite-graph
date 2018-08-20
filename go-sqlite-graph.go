package sqlitegraph

import (
	"errors"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // Database driver
)

//Graph holds the data of the graph with all it's Nodes and edges.
type Graph struct {
	Root  *Node
	Nodes []*Node
	Edges []*Edge
}

//NewGraph creates a new graph containig only the root Node
func NewGraph() *Graph {
	g := new(Graph)

	//Root Node should always be present and have the id 0
	r := NewNode(0)
	g.Root = r
	g.AddNode(r)
	return g
}

//FindEdgesFromTo finds an endge in the graph by it's start and end
func (g *Graph) FindEdgesFromTo(IDFrom, IDTo int) []*Edge {
	var edges []*Edge

	for _, v := range g.Edges {
		if v.From == IDFrom && v.To == IDTo {
			edges = append(edges, v)
		}
	}
	return edges
}

//FindEdgeByID finds a Edge in the graph by it's ID
func (g *Graph) FindEdgeByID(ID int) (*Edge, error) {
	for _, v := range g.Edges {
		if v.ID == ID {
			return v, nil
		}
	}
	return nil, errors.New("Node not found")
}

//FindNodeByID finds a node in the graph by it's ID
func (g *Graph) FindNodeByID(ID int) (*Node, error) {
	for _, v := range g.Nodes {
		if v.ID == ID {
			return v, nil
		}
	}
	return nil, errors.New("Node not found")
}

//ChildsOf finds the childs of a node
func (g *Graph) ChildsOf(n Node) []*Node {

	nodes := []*Node{}

	for _, v := range g.Edges {

		tmpEdge := *v

		if tmpEdge.From == n.ID {
			tmpNode, err := g.FindNodeByID(tmpEdge.To)
			if err != nil {
				panic(err)
			}
			nodes = append(nodes, tmpNode)
		}
	}
	return nodes
}

//ParentsOf finds the parents of a node
func (g *Graph) ParentsOf(ID int) []int {
	nodes := []int{}

	for _, v := range g.Edges {

		tmpEdge := *v

		if tmpEdge.To == ID {
			nodes = append(nodes, tmpEdge.From)
		}
	}
	return nodes
}

// Empty returns true if the root Node is the only Node in the graph, false otherwise
func (g *Graph) Empty() bool {
	return len(g.Nodes) <= 1
}

// AddNode adds a Node to the graph
func (g *Graph) AddNode(n *Node) error {

	//Check if ID already exists in graph
	for _, v := range g.Nodes {
		if v.ID == n.ID {
			return errors.New("go-sqlite-graph: Node ID " + strconv.Itoa(n.ID) + " already exists in graph")
		}
	}

	g.Nodes = append(g.Nodes, n)
	return nil
}

// AddEdge adds a Node to the graph
func (g *Graph) AddEdge(e *Edge) error {

	//Check if ID already exists in graph
	for _, v := range g.Edges {
		if v.ID == e.ID {
			return errors.New("go-sqlite-graph: edge ID " + strconv.Itoa(e.ID) + " already exists in graph")
		}
	}

	g.Edges = append(g.Edges, e)
	return nil

}

//DeleteNode deletes an Node from the graph if it is present
func (g *Graph) DeleteNode(id int) bool {

	for _, v := range g.Nodes {
		if v.ID == id {
			//TODO Delete
			return true
		}
	}
	return false
}

//DeleteEdge deletes an edge from the graph if it is present
func (g *Graph) DeleteEdge(id int) bool {

	for _, v := range g.Edges {
		if v.ID == id {
			//TODO Delete
			return true
		}
	}
	return false
}

// FindSubGraph returns a subset of the graph (a subgraph) with the shortest way
// from the start nodes to the end nodes if possible. If no possible connection is found, an error is returned
func (g *Graph) FindSubGraph(startIDs, endIDs []int) (*Graph, error) {

	out := NewGraph()

	//Iterate over end nodes (multiple recipes are possible)
	for _, e := range endIDs {

		for _, v := range g.findWay(startIDs, e) {

			node, err := g.FindNodeByID(v)
			if err != nil {
				panic(err)
			}

			out.AddNode(node)
		}
	}

	for _, n := range out.Nodes {
		for _, c := range g.ChildsOf(*n) {
			for e := range g.FindEdgesFromTo(n.ID, c.ID) {
				edge, err := g.FindEdgeByID(e)
				if err != nil {
					panic(err)
				}
				g.AddEdge(edge)
			}
		}
	}
	return out, nil
}
