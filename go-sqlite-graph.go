package sqlitegraph

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // Database driver
)

func mainBak() {
	path := "./data1.db"
	os.Remove(path)

	g := NewGraph()
	fmt.Println(g.Empty())

	for i := 1; i < 10; i++ {
		n := NewNode(i)
		err := g.AddNode(n)
		must(err)
	}

	e := NewEdge(0, 0, 1)
	err := g.AddEdge(e)
	must(err)

	e = NewEdge(1, 0, 2)
	err = g.AddEdge(e)
	must(err)

	e = NewEdge(2, 1, 2)
	err = g.AddEdge(e)
	must(err)

	e = NewEdge(3, 2, 3)
	err = g.AddEdge(e)
	must(err)

	fmt.Println("Nodes: " + strconv.Itoa(len(g.Nodes)))
	fmt.Println("Edges: " + strconv.Itoa(len(g.Edges)))

	err = g.Save(path)
	must(err)
	g.PrintGraphviz()

}

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

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// FindSubgraph returns a subset of the graph (a subgraph) with the shortest way
// from the start nodes to the end nodes if possible. If no possible connection is found, an error is returned
func (g *Graph) FindSubgraph(starts, ends []Node) (*Graph, error) {
	sub := NewGraph()
	return sub, nil
}

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
func (g *Graph) ParentsOf(n *Node) []*Node {
	nodes := []*Node{}

	for _, v := range g.Edges {

		tmpEdge := *v

		if tmpEdge.To == n.ID {
			tmpNode, err := g.FindNodeByID(tmpEdge.From)
			if err != nil {
				panic(err)
			}
			nodes = append(nodes, tmpNode)
		}
	}
	return nodes
}

// Save saves the graph to a sqlite database specified by the path
func (g *Graph) Save(path string) error {
	database, err := sql.Open("sqlite3", path)
	must(err)

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS graphNodes (id INTEGER, data TEXT)")
	must(err)
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS graphedges (id INTEGER, nfrom INTEGER, nto INTEGER)")
	must(err)
	statement.Exec()

	statementInsertNodes, err := database.Prepare("INSERT INTO graphNodes (id, data) VALUES (?, ?)")
	must(err)

	statementInsertEdges, err := database.Prepare("INSERT INTO graphedges (id, nfrom, nto ) VALUES (?, ?, ?)")
	must(err)

	fmt.Println("Found " + strconv.Itoa(len(g.Nodes)) + " Nodes")
	fmt.Println("Found " + strconv.Itoa(len(g.Edges)) + " edges")

	for _, v := range g.Nodes {
		fmt.Println("Saving Node: " + strconv.Itoa(v.ID) + " " + v.Text)
		statementInsertNodes.Exec(v.ID, v.Text)
	}
	for _, v := range g.Edges {
		fmt.Println("Saving Edge: " + strconv.Itoa(v.ID) + " " + strconv.Itoa(v.From) + " -> " + strconv.Itoa(v.To))
		statementInsertEdges.Exec(v.ID, v.From, v.To)
	}

	return nil
}

// Load loads a graph from a sqlite database specified by the path
func (g *Graph) Load(path string) error {
	g = NewGraph()
	database, err := sql.Open("sqlite3", path)
	rows, err := database.Query("SELECT id, text FROM graph-Nodes")
	must(err)

	//Load Nodes
	var idNode int
	var text string
	for rows.Next() {
		rows.Scan(&idNode, &text)
		n := NewNode(idNode)
		n.Text = text
		g.AddNode(n)
		fmt.Println(strconv.Itoa(idNode) + ": " + text)
	}

	//Load edges
	var idEdge int
	var from int
	var to int

	for rows.Next() {
		rows.Scan(&idEdge, &from, &to)
		e := NewEdge(idEdge, from, to)
		g.AddEdge(e)
		fmt.Println(strconv.Itoa(idEdge) + ": " + strconv.Itoa(from) + " -> " + strconv.Itoa(to))
	}
	return nil
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

// SubGraph returns a graph containing all nodes necessary to get from the starts to the ends
// It will return an error, if no way is found
func (g *Graph) SubGraph(startIDs, endIDs []int) (*Graph, error) {
	//TODO implement
	return g, nil
}

//PrintGraphviz generates a graph in the dot language to be visualized using graphviz
func (g *Graph) PrintGraphviz() error {
	fmt.Println("digraph {")

	for _, v := range g.Nodes {

		fmt.Println("	" + strconv.Itoa(v.ID) + " [label=\"ID: " + strconv.Itoa(v.ID) + " DATA: " + v.Text + "\"];")
	}

	for _, v := range g.Edges {
		fmt.Println("	" + strconv.Itoa(v.From) + " -> " + strconv.Itoa(v.To) + ";")
	}
	fmt.Println("}")

	return nil
}
