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

	g4 := NewGraph()
	g4n1 := NewNode(1)
	g4n2 := NewNode(2)
	g4n3 := NewNode(3)
	g4n4 := NewNode(4)
	g4n5 := NewNode(5)
	g4n6 := NewNode(6)

	g4.AddNode(g4n2)
	g4.AddNode(g4n6)
	g4.AddNode(g4n3)
	g4.AddNode(g4n4)
	g4.AddNode(g4n5)

	g4Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{
			g4n1,
			g4n2,
			g4n3,
			g4n4,
			g4n5,
			g4n6,
	},
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
		{"Check if nodes are sorted", g4, g4n1, g4Result, false},
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
	e1 := NewEdge(1, 1, 2)
	g1 := NewGraph()
	g1Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{},
		Edges: []*Edge{e1},
	}

	// g2
	e2 := NewEdge(1, 1, 2)
	e2Pre := NewEdge(1, 2, 3)
	g2 := NewGraph()
	g2.AddEdge(e2Pre)
	g2Result := &Graph{
		Root:  NewNode(0),
		Nodes: []*Node{},
		Edges: []*Edge{e2Pre},
	}

	// g3
	e3 := NewEdge(1, 1, 2)
	e4 := NewEdge(2, 1, 2)
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

func TestNewGraph(t *testing.T) {
	n1 := NewNode(0)
	g1 := &Graph{
		Root:  n1,
		Nodes: []*Node{n1},
		Edges: []*Edge{},
	}
	tests := []struct {
		name string
		want *Graph
	}{
		{"Create new empty graph", g1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(); g1.Root.ID != 0 {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindEdgesFromTo(t *testing.T) {

	e1 := NewEdge(1, 1, 2)
	e2 := NewEdge(2, 2, 3)
	e3 := NewEdge(3, 1, 2)

	g0 := NewGraph()
	e0Result := []*Edge{}

	g1 := NewGraph()
	g1.AddEdge(e1)
	g1.AddEdge(e2)

	e1Result := []*Edge{e1}

	g2 := NewGraph()
	g2.AddEdge(e1)
	g2.AddEdge(e2)
	g2.AddEdge(e3)
	e2Result := []*Edge{e1, e3}

	type args struct {
		IDFrom int
		IDTo   int
	}
	tests := []struct {
		name  string
		graph *Graph
		args  args
		want  []*Edge
	}{
		{"Find no edges", g0, args{1, 2}, e0Result},
		{"Find one edge", g1, args{1, 2}, e1Result},
		{"Find two edges", g2, args{1, 2}, e2Result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			foundEdges := g.FindEdgesFromTo(tt.args.IDFrom, tt.args.IDTo)

			if len(foundEdges) != len(tt.want) {
				t.Errorf("len(Graph.FindEdgesFromTo()) = %v, want %v", len(foundEdges), len(tt.want))
			}

			for k, v := range foundEdges {
				if v.ID != tt.want[k].ID {
					t.Errorf("Graph.FindEdgesFromTo() = %v, want %v", foundEdges, tt.want)
				}
			}
		})
	}
}

func TestGraph_FindEdgeByID(t *testing.T) {

	g1 := NewGraph()
	g2 := NewGraph()
	g3 := NewGraph()

	e1 := NewEdge(1, 1, 2)
	e2 := NewEdge(2, 1, 2)
	e3 := NewEdge(3, 1, 2)

	g1.AddEdge(e1)
	g1.AddEdge(e2)
	g1.AddEdge(e3)

	g2.AddEdge(e1)
	g2.AddEdge(e2)

	tests := []struct {
		name    string
		graph   *Graph
		ID      int
		want    *Edge
		wantErr bool
	}{
		{"Look for edge that exists", g1, 3, e3, false},
		{"Look for edge that does not exist", g2, 3, nil, true},
		{"Search empty graph", g3, 3, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			got, err := g.FindEdgeByID(tt.ID)
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

	g1 := NewGraph()
	g2 := NewGraph()
	g3 := NewGraph()
	g4 := NewGraph()
	g5 := NewGraph()

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	g1.AddNode(n1)
	g1.AddNode(n2)
	g1.AddNode(n3)

	g2.AddNode(n1)
	g2.AddNode(n2)

	g5.AddNode(n1)
	g5.AddNode(n2)
	g5.AddNode(n3)

	tests := []struct {
		name    string
		graph   *Graph
		ID      int
		want    *Node
		wantErr bool
	}{
		{"Look for edge that exists", g1, 3, n3, false},
		{"Look for edge that does not exist", g2, 3, nil, true},
		{"Search empty graph", g3, 3, nil, true},
		{"Search for root node in empty graph", g4, 0, g4.Root, false},
		{"Search for root node in graph with nodes", g5, 0, g5.Root, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			got, err := g.FindNodeByID(tt.ID)
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

func TestGraph_FindSubGraph(t *testing.T) {

	g1 := NewGraph()

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)
	n7 := NewNode(7)
	n8 := NewNode(8)
	n9 := NewNode(9)
	n10 := NewNode(10)

	e1 := NewEdge(1, 0, 1)
	e2 := NewEdge(2, 0, 4)
	e3 := NewEdge(3, 0, 2)
	e4 := NewEdge(4, 0, 3)
	e5 := NewEdge(5, 4, 5)
	e6 := NewEdge(6, 6, 5)
	e7 := NewEdge(7, 6, 9)
	e8 := NewEdge(8, 7, 9)
	e9 := NewEdge(9, 2, 8)
	e10 := NewEdge(10, 2, 7)
	e11 := NewEdge(11, 1, 2)
	e12 := NewEdge(12, 8, 10)
	e13 := NewEdge(13, 3, 6)

	nodes := []*Node{n1, n2, n3, n4, n5, n6, n7, n8, n9, n10}
	edges := []*Edge{e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13}

	for _, v := range nodes {
		g1.AddNode(v)
	}

	for _, v := range edges {
		g1.AddEdge(v)
	}

	g1Result := NewGraph()

	g1Result.AddNode(n3)
	g1Result.AddNode(n6)
	g1Result.AddNode(n7)
	g1Result.AddNode(n9)

	g1Result.AddEdge(e7)
	g1Result.AddEdge(e8)
	g1Result.AddEdge(e13)

	g1StartIDs := []int{7, 3}
	g1EndIDs := []int{9}

	type args struct {
		startIDs []int
		endIDs   []int
	}
	tests := []struct {
		name    string
		graph   *Graph
		args    args
		want    *Graph
		wantErr bool
	}{
		{"Test1", g1, args{g1StartIDs, g1EndIDs}, g1Result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			got, err := g.FindSubGraph(tt.args.startIDs, tt.args.endIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.FindSubGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.want.Equal(got) {
				t.Errorf("Graph.FindSubGraph() = \n%v, want \n%v", got.ToString(), tt.want.ToString())
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
		name    string
		fields  fields
		args    args
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
			if err := g.DeleteNode(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Graph.DeleteNode() error = %v, wantErr %v", err, tt.wantErr)
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
		name    string
		fields  fields
		args    args
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
			if err := g.DeleteEdge(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Graph.DeleteEdge() error = %v, wantErr %v", err, tt.wantErr)
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
		n int
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
			if got := g.ChildsOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.ChildsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Equal(t *testing.T) {
	n1 :=NewNode(1)
	n2 :=NewNode(2)
	e1 :=NewEdge(1,1,2)
	e2 :=NewEdge(2,0,2)

	//identical filled graphs
	g1a := NewGraph()
	g1b := NewGraph()

	g1a.AddNode(n1)
	g1a.AddNode(n2)
	g1a.AddEdge(e1)
	g1a.AddEdge(e2)

	g1b.AddNode(n1)
	g1b.AddNode(n2)
	g1b.AddEdge(e1)
	g1b.AddEdge(e2)

	// identical empty graphs
	g2a := NewGraph()
	g2b := NewGraph()

	//node length unewqeal
	g3a := NewGraph()
	g3b := NewGraph()

	g3a.AddNode(n1)
	g3a.AddNode(n2)
	g3a.AddEdge(e1)
	g3a.AddEdge(e2)

	g3b.AddNode(n2)
	g3b.AddEdge(e1)
	g3b.AddEdge(e2)

	//edeg length unequal
	g4a := NewGraph()
	g4b := NewGraph()

	g4a.AddNode(n1)
	g4a.AddNode(n2)
	g4a.AddEdge(e1)

	g4b.AddNode(n1)
	g4b.AddNode(n2)
	g4b.AddEdge(e1)
	g4b.AddEdge(e2)

	//different node id
	g5a := NewGraph()
	g5b := NewGraph()

	g5a.AddEdge(e1)
	g5a.AddNode(n1)

	g5b.AddEdge(e1)
	g5b.AddNode(n2)

	//different edge id
	g6a := NewGraph()
	g6b := NewGraph()

	g6a.AddEdge(e1)
	g6a.AddNode(n1)
	g6a.AddNode(n2)

	g6b.AddEdge(e2)
	g6b.AddNode(n1)
	g6b.AddNode(n2)

	tests := []struct {
		name   string
		ga *Graph
		gb *Graph
		want   bool
	}{
		{"Identical graphs (filled)",g1a, g1b, true},
		{"Identical graphs (empty)", g2a, g2b, true},
		{"Node Length unequal", g3a, g3b, false},
		{"Edge Length unequal", g4a, g4b, false},
		{"Different Node ID", g5a, g5b, false},
		{"Different Edge ID", g6a, g6b, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ga.Equal(tt.gb); got != tt.want {
				t.Errorf("Graph.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_ParentsOf(t *testing.T) {

	g1 := NewGraph()
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)
	n7 := NewNode(7)
	n8 := NewNode(8)

	g1.AddNode(n1)
	g1.AddNode(n2)
	g1.AddNode(n3)
	g1.AddNode(n4)
	g1.AddNode(n5)
	g1.AddNode(n6)
	g1.AddNode(n7)
	g1.AddNode(n8)

	g1.AddEdge(NewEdge(1, 1,2))

	g2 := NewGraph()
	g2.AddNode(n1)
	g2.AddNode(n2)
	g2.AddNode(n3)
	g2.AddNode(n4)
	g2.AddNode(n5)
	g2.AddNode(n6)
	g2.AddNode(n7)
	g2.AddNode(n8)

	g2.AddEdge(NewEdge(1,1,2))
	g2.AddEdge(NewEdge(2,3,2))
	g2.AddEdge(NewEdge(3,4,2))
	g2.AddEdge(NewEdge(4,5,2))


	tests := []struct {
		name   string
		graph *Graph
		node int
		want   []int
	}{
		{"Find one parent", g1, 2, []int{1}},
		{"Find four parents", g2, 2, []int{1,3,4,5}},
		{"Find no parents", g1, 8, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			if got := g.ParentsOf(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.ParentsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
