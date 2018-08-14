package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path := "./data1.db"
	g := NewGraph()
	fmt.Println(g.Empty())
	for i := 1; i < 10; i++ {
		n := NewNode(i)
		g.AddNode(n)
	}
	fmt.Println(len(g.Nodes))
	fmt.Println(g.Empty())
	g.Save(path)
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
	g.Root = NewNode(0)
	return g
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Save saves the graph to a sqlite database specified by the path
func (g *Graph) Save(path string) {
	fmt.Println("Found " + strconv.Itoa(len(g.Nodes)) + " nodes")
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

	for k, v := range g.Nodes {
		statementInsertNodes.Exec(k, v.Text)
		fmt.Println("Saving Node: " + strconv.Itoa(v.Id) + " " + v.Text)
		for k2, v2 := range v.Children {
			statementInsertEdges.Exec(k2, v.Id, v2.Id)
			fmt.Println("Saving Edge: " + strconv.Itoa(v2.Id) + " " + strconv.Itoa(v2.Id) + " -> " + strconv.Itoa(v2.Id))
		}
	}

}

//DeleteGraphFromDB deletes all nodes and edges belonging to a certain graph id in the database
func DeleteGraphFromDB(id int) bool {
	//TODO implement
	return false
}

//TODO add graph id to database

// Load loads a graph from a sqlite database specified by the path
func (g *Graph) Load(path string) {
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
		e := NewEdge(from, to)
		e.Id = idEdge
		g.AddEdge(e)
		fmt.Println(strconv.Itoa(idEdge) + ": " + strconv.Itoa(from) + " -> " + strconv.Itoa(to))
	}
}

// Empty returns true if the root node is the only node in the graph, false otherwise
func (g *Graph) Empty() bool {
	return len(g.Nodes) <= 1
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) bool {

	//Check if Id already exists in graph
	for _, v := range g.Nodes {
		if v.Id == n.Id {
			return false
		}
	}

	g.Nodes = append(g.Nodes, n)
	return true
}

// AddNode adds a node to the graph
func (g *Graph) AddEdge(e *Edge) bool {

	//Check if Id already exists in graph
	for _, v := range g.Edges {
		if v.Id == e.Id {
			return false
		}
	}

	g.Edges = append(g.Edges, e)
	return true

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
