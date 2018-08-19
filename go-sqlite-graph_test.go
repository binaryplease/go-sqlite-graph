package sqlitegraph

import (
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

	for n := range nodes { g.AddNode(nodes[n]) }
	for e := range edges { g.AddEdge(edges[e]) }
	return g
}

func TestGraph_Save(t *testing.T) {
	g := createTestGraph()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		graph   *Graph
		args    args
		wantErr bool
	}{
		{"Should succeed", g, args{path: "./testdb.db"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := g.Save(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Graph.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraph_Load(t *testing.T) {

	g := createTestGraph()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
        g       *Graph
		args    args
		wantErr bool
	}{
		{"Should succeed", g, args{path: "testload1.db"}, false},
		{"Should fail", g, args{path: "testload2.db"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if err := g.Load(tt.args.path); (err != nil) != tt.wantErr {
			// 	t.Errorf("Graph.Load() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}

func TestGraph_Empty(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	tests := []struct {
		name   string
		fields fields
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
			if got := g.Empty(); got != tt.want {
				t.Errorf("Graph.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddNode(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		n *Node
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
			if err := g.AddNode(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		e *Edge
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
			if err := g.AddEdge(tt.args.e); (err != nil) != tt.wantErr {
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

func TestGraph_PrintGraphviz(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := g.PrintGraphviz(); (err != nil) != tt.wantErr {
				t.Errorf("Graph.PrintGraphviz() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
