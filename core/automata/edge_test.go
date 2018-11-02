package automata

import (
	"reflect"
	"testing"
)

func TestEdge_copy(t *testing.T) {
	tests := []struct {
		name string
		e    Edge
		want Edge
	}{
		{"Copies edge", Edge{From: 0, To: 1, Label: 'a'}, Edge{From: 0, To: 1, Label: 'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edge.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEdge_copyWithOffset(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name string
		e    Edge
		args args
		want Edge
	}{
		{"Doesn't offset initial", Edge{From: 0, To: 1, Label: 'a'}, args{offset: 2}, Edge{From: 0, To: 3, Label: 'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.copyWithOffset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edge.copyWithOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
