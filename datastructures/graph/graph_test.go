package graph

import (
	"testing"
)

var (
	testGraph0 = Graph{}
	testGraph1 = Graph{nodes: []*GraphNode{&GraphNode{value: 1, nodes: nil}}}
	testGraph2 = Graph{
		nodes: []*GraphNode{
			&GraphNode{value: 1, nodes: []*GraphNode{
				&GraphNode{value: 2, nodes: nil},
				&GraphNode{value: 3, nodes: nil},
				&GraphNode{value: 4, nodes: nil},
			},
			},
		},
	}
	testGraph3 = Graph{
		nodes: []*GraphNode{
			&GraphNode{value: 1, nodes: []*GraphNode{
				&GraphNode{value: 2, nodes: nil},
				&GraphNode{value: 3, nodes: nil},
				&GraphNode{value: 4, nodes: []*GraphNode{
					&GraphNode{value: 5, nodes: []*GraphNode{
						&GraphNode{value: 6, nodes: nil},
						},
						},
					},
				},
				},
			},
		},
	}
)

func TestDFS(t *testing.T) {
	tests := []struct {
		name string
		graph Graph
		want []*GraphNode
	}{
		{"empty", testGraph0, nil},
		{"size-1", testGraph1, []*GraphNode{{value: 1}}},
		{"size-4", testGraph2, []*GraphNode{&GraphNode{value: 4, nodes: nil},&GraphNode{value: 3, nodes: nil},&GraphNode{value: 2, nodes: nil},&GraphNode{value: 1, nodes: nil}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.graph.DFS()
			if len(got) != len(tt.want) {
				t.Errorf("wantSize: %d | gotSize: %d", len(tt.want), len(got))
			}

			for i := range got {
				if got[i].value != tt.want[i].value {
					t.Errorf("got: %#v\nwant: %#v", got[i].value, tt.want[i].value)
				}
			}
		})
	}
}