package sqlitegraph

import (
	"reflect"
	"testing"
)

func TestNewEdge(t *testing.T) {
	type args struct {
		id   int
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want *Edge
	}{
		{"Create simple Edge", args{0, 1, 2}, &Edge{ID: 0, From: 1, To: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEdge(tt.args.id, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}
