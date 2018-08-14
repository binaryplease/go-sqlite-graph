package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
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

	fmt.Println("Nodes: " + strconv.Itoa(len(g.Nodes)))
	fmt.Println("Edges: " + strconv.Itoa(len(g.Edges)))

	err = g.Save(path)
	must(err)
	g.PrintGraphviz()

}

//Graph holds the data of the graph with all it's nodes and edges.
type Graph struct {
	Root  *Node
	Nodes []*Node
	Edges []*Edge
}

//NewGraph creates a new graph containig only the root node
func NewGraph() *Graph {
	g := new(Graph)

	//Root node should always be present and have the id 0
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

// Save saves the graph to a sqlite database specified by the path
func (g *Graph) Save(path string) error {
	database, err := sql.Open("sqlite3", path)
	must(err)

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS graphnodes (id INTEGER, data TEXT)")
	must(err)
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS graphedges (id INTEGER, nfrom INTEGER, nto INTEGER)")
	must(err)
	statement.Exec()

	statementInsertNodes, err := database.Prepare("INSERT INTO graphnodes (id, data) VALUES (?, ?)")
	must(err)

	statementInsertEdges, err := database.Prepare("INSERT INTO graphedges (id, nfrom, nto ) VALUES (?, ?, ?)")
	must(err)

	fmt.Println("Found " + strconv.Itoa(len(g.Nodes)) + " nodes")
	fmt.Println("Found " + strconv.Itoa(len(g.Edges)) + " edges")

	for _, v := range g.Nodes {
		fmt.Println("Saving Node: " + strconv.Itoa(v.Id) + " " + v.Text)
		statementInsertNodes.Exec(v.Id, v.Text)
	}
	for _, v := range g.Edges {
		fmt.Println("Saving Edge: " + strconv.Itoa(v.Id) + " " + strconv.Itoa(v.From) + " -> " + strconv.Itoa(v.To))
		statementInsertEdges.Exec(v.Id, v.From, v.To)
	}

	return nil
}

// Load loads a graph from a sqlite database specified by the path
func (g *Graph) Load(path string) error {
	g = NewGraph()
	database, err := sql.Open("sqlite3", path)
	rows, err := database.Query("SELECT id, text FROM graph-nodes")
	must(err)

	//Load nodes
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

// Empty returns true if the root node is the only node in the graph, false otherwise
func (g *Graph) Empty() bool {
	return len(g.Nodes) <= 1
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) error {

	//Check if Id already exists in graph
	for _, v := range g.Nodes {
		if v.Id == n.Id {
			return errors.New("go-sqlite-graph: node ID " + strconv.Itoa(n.Id) + " already exists in graph")
		}
	}

	g.Nodes = append(g.Nodes, n)
	return nil
}

// AddNode adds a node to the graph
func (g *Graph) AddEdge(e *Edge) error {

	//Check if Id already exists in graph
	for _, v := range g.Edges {
		if v.Id == e.Id {
			return errors.New("go-sqlite-graph: edge ID " + strconv.Itoa(e.Id) + " already exists in graph")
		}
	}

	g.Edges = append(g.Edges, e)
	return nil

}

//Deletenode deletes an node from the graph if it is present
func (g *Graph) DeleteNode(id int) bool {

	for _, v := range g.Nodes {
		if v.Id == id {
			//TODO Delete
			return true
		}
	}
	return false
}

//DeleteEdge deletes an edge from the graph if it is present
func (g *Graph) DeleteEdge(id int) bool {

	for _, v := range g.Edges {
		if v.Id == id {
			//TODO Delete
			return true
		}
	}
	return false
}

//PrintGraphviz generates a graph in the dot language to be visualized using graphviz
func (g *Graph) PrintGraphviz() error {
	fmt.Println(" digraph {")

	for _, v := range g.Nodes {

		fmt.Println("	" + strconv.Itoa(v.Id) + " [label=\"ID: " + strconv.Itoa(v.Id) + " DATA: " + v.Text + "\"];")
	}

	for _, v := range g.Edges {
		fmt.Println("	" + strconv.Itoa(v.From) + " -> " + strconv.Itoa(v.To) + ";")
	}
	fmt.Println("}")

	return nil
}
