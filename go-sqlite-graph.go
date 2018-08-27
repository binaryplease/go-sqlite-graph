package sqlitegraph

import (
	"errors"
	"sort"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // Database driver
)

//Graph holds the data of the graph with all it's Nodes and edges.
type Graph struct {
	Root  *Node
	Nodes []*Node
	Edges []*Edge
}

func (g Graph) ToString() string {
	result := "Nodes: ["

	for _, v := range g.Nodes {
		result = result + "(ID: " + strconv.Itoa(v.ID) + ") "
	}
	result = result + "] Edges: ["

	for _, v := range g.Edges {
		result = result + "(ID: " + strconv.Itoa(v.ID) + ") "
	}

	result = result + "]"
	return result
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
func (g *Graph) ChildsOf(n int) []int {

	nodes := []int{}

	for _, v := range g.Edges {

		tmpEdge := *v

		if tmpEdge.From == n {
			tmpNode := (tmpEdge.To)
			nodes = append(nodes, tmpNode)
		}
	}
	return nodes
}

//Check if two graphs are the same
func (g Graph) Equal(g2 *Graph) bool {

	if len(g.Nodes) != len(g2.Nodes) {
		return false
	}

	if len(g.Edges) != len(g2.Edges) {
		return false
	}

	//Compare nodes
	for k, _ := range g.Nodes {
		if !g.Nodes[k].Equals(g2.Nodes[k]) {
			return false
		}
	}

	//Compare Edges
	for k, _ := range g.Edges {
		if !g.Edges[k].Equals(g2.Edges[k]) {
			return false
		}
	}

	return true
}

//ParentsOf finds the parents of a node
func (g *Graph) ParentsOf(n int) []int {
	nodes := []int{}

	for _, v := range g.Edges {


		if v.To == n {
			nodes = append(nodes, v.From)
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

	//TODO sort
	//Check if ID already exists in graph
	for _, v := range g.Nodes {
		if v.ID == n.ID {
			return errors.New("go-sqlite-graph: Node ID " + strconv.Itoa(n.ID) + " already exists in graph")
		}
	}

	g.Nodes = append(g.Nodes, n)
	sort.Slice(g.Nodes, func(i, j int) bool {
		return g.Nodes[i].ID < g.Nodes[j].ID
	})

	return nil
}

// AddEdge adds a Node to the graph
func (g *Graph) AddEdge(e *Edge) error {
	//TODO sort

	//Check if ID already exists in graph
	for _, v := range g.Edges {
		if v.ID == e.ID {
			return errors.New("go-sqlite-graph: edge ID " + strconv.Itoa(e.ID) + " already exists in graph")
		}
	}

	g.Edges = append(g.Edges, e)
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].ID < g.Edges[j].ID
	})
	return nil

}

//DeleteNode deletes an Node from the graph if it is present
func (g *Graph) DeleteNode(id int) error {

	for _, v := range g.Nodes {
		if v.ID == id {
			//TODO Delete
			return nil
		}
	}
	return errors.New("Could not find Node with ID " + strconv.Itoa(id) + " in graph")
}

//DeleteEdge deletes an edge from the graph if it is present
func (g *Graph) DeleteEdge(id int) error {

	for _, v := range g.Edges {
		if v.ID == id {
			//TODO Delete
			return nil
		}
	}
	return errors.New("Could not find Edge with ID " + strconv.Itoa(id) + " in graph")
}

// FindSubGraph returns a subset of the graph (a subgraph) with the shortest way
// from the start nodes to the end nodes if possible. If no possible connection is found, an error is returned
func (g *Graph) FindSubGraph(startIDs, endIDs []int) (*Graph, error) {

	out := NewGraph()

	//Iterate over end nodes (multiple recipes are possible)
	for _, e := range endIDs {
		n, err := g.FindNodeByID(e)
		checkErr(err)
		out.AddNode(n)

		if contains(startIDs, e) {
			continue
		} else {
			for _,c := range g.ParentsOf(e) {
				edge  := g.FindEdgesFromTo(c,e)[0]
				out.AddEdge(edge)


				sub, err := g.FindSubGraph(startIDs, []int{c})
				checkErr(err)
				for _,v := range sub.Edges{
					out.AddEdge(v)
				}
				for _,v := range sub.Nodes {
					out.AddNode(v)
				}
			}
		}
	}
return out, nil
}

