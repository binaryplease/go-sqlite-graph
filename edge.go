package sqlitegraph

//Edge is the data structure to connect nodes inside the graph.
type Edge struct {
	ID   int
	From int
	To   int
	Text string
}

// NewEdge creates and returns a new edge from a given start and end node id
func NewEdge(id, from, to int) *Edge {
	e := new(Edge)
	e.ID = id
	e.From = from
	e.To = to
	return e
}
