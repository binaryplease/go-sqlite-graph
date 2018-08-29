package sqlitegraph

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Save saves the graph to a sqlite database specified by the path
func (g *Graph) Save(path string) error {
	database, err := sql.Open("sqlite3", path)
	checkErr(err)

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS graphNodes (id INTEGER, data TEXT)")
	checkErr(err)
	statement.Exec()

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS graphedges (id INTEGER, nfrom INTEGER, nto INTEGER)")
	checkErr(err)
	statement.Exec()

	statementInsertNodes, err := database.Prepare("INSERT INTO graphNodes (id, data) VALUES (?, ?)")
	checkErr(err)

	statementInsertEdges, err := database.Prepare("INSERT INTO graphedges (id, nfrom, nto ) VALUES (?, ?, ?)")
	checkErr(err)

	for _, v := range g.Nodes {
		// fmt.Println("Saving Node: " + strconv.Itoa(v.ID) + " " + v.Text)
		statementInsertNodes.Exec(v.ID, v.Text)
	}
	for _, v := range g.Edges {
		// fmt.Println("Saving Edge: " + strconv.Itoa(v.ID) + " " + strconv.Itoa(v.From) + " -> " + strconv.Itoa(v.To))
		statementInsertEdges.Exec(v.ID, v.From, v.To)
	}

	return nil
}

// Load loads a graph from a sqlite database specified by the path
func (g *Graph) Load(path string) error {
	g = NewGraph()
	database, err := sql.Open("sqlite3", path)
	checkErr(err)
	rows, err := database.Query("SELECT id, data FROM graphNodes")
	checkErr(err)

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
