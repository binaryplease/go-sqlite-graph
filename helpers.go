package sqlitegraph

import (
	"fmt"
	"os"
	"strconv"
)

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

//Check errors, TODO do something better
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func mainBak() {
	path := "./data1.db"
	os.Remove(path)

	g := NewGraph()
	fmt.Println(g.Empty())

	for i := 1; i < 10; i++ {
		n := NewNode(i)
		err := g.AddNode(n)
		checkErr(err)
	}

	e := NewEdge(0, 0, 1)
	err := g.AddEdge(e)
	checkErr(err)

	e = NewEdge(1, 0, 2)
	err = g.AddEdge(e)
	checkErr(err)

	e = NewEdge(2, 1, 2)
	err = g.AddEdge(e)
	checkErr(err)

	e = NewEdge(3, 2, 3)
	err = g.AddEdge(e)
	checkErr(err)

	fmt.Println("Nodes: " + strconv.Itoa(len(g.Nodes)))
	fmt.Println("Edges: " + strconv.Itoa(len(g.Edges)))

	err = g.Save(path)
	checkErr(err)
	g.PrintGraphviz()

}

// contains checks if a int is in a given array
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//findWay recursively goes up the graph stopping at given start nodes
func (g *Graph) findWay(startIDs []int, endID int) []int {

	ids := []int{endID}
	ids = append(ids, startIDs...)

	for _, v := range g.ParentsOf(endID) {

		//If parent is not a startID
		if !contains(startIDs, v) {
			// Add parents to list
			ids = append(ids, v)
			ids = append(ids, g.findWay(startIDs, v)...)
		}
	}
	return ids
}
