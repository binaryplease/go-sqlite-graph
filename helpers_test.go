package sqlitegraph

import (
	"reflect"
	"testing"
)

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

func Test_checkErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkErr(tt.args.err)
		})
	}
}

func Test_mainBak(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			mainBak()
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_findWay(t *testing.T) {
	type fields struct {
		Root  *Node
		Nodes []*Node
		Edges []*Edge
	}
	type args struct {
		startIDs []int
		endID    int
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
			if got := g.findWay(tt.args.startIDs, tt.args.endID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.findWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
