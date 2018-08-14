package sqlitegraph

type Edge struct {
	Id int
	From int
	To   int
}

func NewEdge(id, from, to int) *Edge {
	e := new(Edge)
	e.Id = id
	e.From = from
	e.To = to
	return e
}
