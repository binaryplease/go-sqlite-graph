package sqlitegraph

import "testing"

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
