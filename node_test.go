package sqlitegraph

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"Create simple node", args{0}, &Node{ID: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
