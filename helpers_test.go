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
		{"Array contains element one time", args{[]int{1,2,3,4,5}, 1}, true},
		{"Array contains element two times", args{[]int{1,2,2,4,5}, 2}, true},
		{"Array does not contain element", args{[]int{1,3,4,5}, 2}, false},
		{"Array is empty", args{[]int{}, 2}, false},
		{"Array has only given element", args{[]int{2}, 2}, true},
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

