package main

type Edge struct {
	Id int
	From int
	To   int
}

func NewEdge(from, to int) *Edge {
	e := new(Edge)
	e.From = from
	e.To = to
	return e
}
