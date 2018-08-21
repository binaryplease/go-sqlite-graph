package sqlitegraph

import (
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func createTestGraph() *Graph {

	/*
		[ROOT 0]
		   |
		(edge 0)
		   |
		   V
		[NODE 1] -(edge 1)-> [NODE 2] -(edge 2)-> [NODE 3]
							   |
							(edge 3)
							   |
							   V
							[NODE 4] -(edge 4)-> [NODE 5]
							   |                    |
							(edge 5)             (edge 6)
							   |                    |
							   V                    V
							[NODE 6] -(edge 7)-> [NODE 7] -(edge 8)-> [NODE 8]
										            |
										         (edge 9)
										            |
										            V
										         [NODE 9] -(edge 10)-> [NODE 10]
	*/
	g := NewGraph()

	nodes := []*Node{
		NewNode(1),
		NewNode(2),
		NewNode(3),
		NewNode(4),
		NewNode(5),
		NewNode(6),
		NewNode(7),
		NewNode(8),
		NewNode(9),
		NewNode(10)}

	edges := []*Edge{
		NewEdge(0, 0, 1),
		NewEdge(1, 1, 2),
		NewEdge(2, 2, 3),
		NewEdge(3, 2, 4),
		NewEdge(4, 4, 5),
		NewEdge(5, 4, 6),
		NewEdge(6, 4, 7),
		NewEdge(7, 6, 7),
		NewEdge(8, 7, 8),
		NewEdge(9, 7, 9),
		NewEdge(10, 9, 10)}

	for n := range nodes {
		g.AddNode(nodes[n])
	}
	for e := range edges {
		g.AddEdge(edges[e])
	}
	return g
}

func TestGraph_Empty(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)

	g1 := NewGraph()
	g2 := NewGraph()
	g3 := NewGraph()

	g2.AddNode(n1)
	g3.AddNode(n1)
	g3.AddNode(n2)

	tests := []struct {
		name  string
		graph *Graph
		want  bool
	}{
		{"Empty graph", g1, true},
		{"One element", g2, false},
		{"Two elements", g3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			if got := g.Empty(); got != tt.want {
				t.Errorf("Graph.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddNode(t *testing.T) {

	//Empty Graph
	g1 := NewGraph()
	n1 := NewNode(1)

	g1Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{n1},
		Edges: []*Edge{},
	}

	g2 := NewGraph()
	n2 := NewNode(1)
	g2.AddNode(n2)

	g2Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{n2},
		Edges: []*Edge{},
	}

	g3 := NewGraph()
	nPre := NewNode(1)
	n3 := NewNode(2)
	g3.AddNode(nPre)

	g3Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{nPre, n3},
		Edges: []*Edge{},
	}

	type args struct {
		n *Node
	}

	tests := []struct {
		name    string
		graph   *Graph
		n       *Node
		want    *Graph
		wantErr bool
	}{
		{"Add a node to empty graph", g1, n1, g1Result, false},
		{"Try to add node with same ID", g2, n2, g2Result, true},
		{"Add a node to graph with one node", g3, n3, g3Result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			if err := g.AddNode(tt.n); (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {

	// g1
	e1 := NewEdge(1,1,2)
	g1 := NewGraph()
	g1Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{},
		Edges: []*Edge{e1},
	}

	// g2
	e2 := NewEdge(1,1,2)
	e2Pre := NewEdge(1,2,3)
	g2 := NewGraph()
	g2.AddEdge(e2Pre)
	g2Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{},
		Edges: []*Edge{e2Pre},
	}

	// g3
	e3 := NewEdge(1,1,2)
	e4 := NewEdge(2,1,2)
	g3 := NewGraph()
	g3.AddEdge(e4)
	g3Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{},
		Edges: []*Edge{e4, e3},
	}

	tests := []struct {
		name    string
		graph   *Graph
		e       *Edge
		want    *Graph
		wantErr bool
	}{
		{"Add a edge to empty graph", g1, e1, g1Result, false},
		{"Try to add edge with same ID", g2, e2, g2Result, true},
		{"Add a second edge to graph with one edge", g3, e3, g3Result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			if err := g.AddEdge(tt.e); (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddEdge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraph_DeleteNode(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.DeleteNode(tt.args.id); got != tt.want {
				t.Errorf("Graph.DeleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_DeleteEdge(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.DeleteEdge(tt.args.id); got != tt.want {
				t.Errorf("Graph.DeleteEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGraph(t *testing.T) {
	tests := []struct {
		name string
		want *Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindEdgesFromTo(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		IDFrom int
		IDTo   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Edge
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.FindEdgesFromTo(tt.args.IDFrom, tt.args.IDTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindEdgesFromTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindEdgeByID(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Edge
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			got, err := g.FindEdgeByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.FindEdgeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindEdgeByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindNodeByID(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			got, err := g.FindNodeByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.FindNodeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindNodeByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_ChildsOf(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		n Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.ChildsOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.ChildsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_ParentsOf(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.ParentsOf(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.ParentsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindSubGraph(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		startIDs []int
		endIDs   []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Graph
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Root:  tt.fields.Root,
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			got, err := g.FindSubGraph(tt.args.startIDs, tt.args.endIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.FindSubGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindSubGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
