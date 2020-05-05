package binary_search

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t0 := bst{}
	t1 := bst{}
	t1.Insert(2)
	t2 := bst{}
	t2.Insert(2)
	t2.Insert(1)
	t3 := bst{}
	t3.Insert(2)
	t3.Insert(1)
	t3.Insert(4)
	t3.Insert(5)
	t4 := bst{}
	t4.Insert(2)
	t4.Insert(1)
	t4.Insert(4)
	t4.Insert(5)
	t4.Insert(0)
	t4.Insert(3)
	t4.Insert(6)
	t4.Insert(7)

	tests := []struct {
		name string
		tree bst
		want []*node
	}{
		{"empty", t0, nil},
		{"size-1", t1, []*node{{data: 2}}},
		{"size-2", t2, []*node{{data: 1}, {data: 2}}},
		{"size-4", t3, []*node{{data: 1}, {data: 2}, {data: 4}, {data: 5}}},
		{"size-8", t4, []*node{{data: 0},{data: 1},{data: 2},{data: 3}, {data: 4}, {data: 5}, {data: 6},{data: 7}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.InorderDFS()
			if len(got) != len(tt.want) {
				t.Errorf("Array lengths not equal:\n\twant:%d\n\tgot: %d", len(tt.want), len(got))
			}

			if got != nil {
				for i := range got {
					if got[i].data !=tt.want[i].data {
						t.Errorf("Array item not equal:\n\twant:%#v\n\tgot: %#v\n\t:index: %d", tt.want[i].data, got[i].data, i)
					}
				}
			}
		})
	}
}

func TestFind(t *testing.T) {
	t0 := bst{}
	t1 := bst{}
	t1.Insert(2)
	t2 := bst{}
	t2.Insert(2)
	t2.Insert(1)
	t3 := bst{}
	t3.Insert(2)
	t3.Insert(1)
	t3.Insert(4)
	t3.Insert(5)

	tests := []struct {
		name string
		tree bst
		find int
		want *node
		wantParent *node
	}{
		{"empty", t0, 3, nil, nil},
		{"size-1", t1, 3, nil, nil},
		{"size-1 | item present", t1, 2, &node{data: 2}, nil},
		{"size-2 | item present", t2, 1, &node{data: 1},  &node{data: 2}},
		{"size-2 | item NOT present", t2, 3, nil, nil},
		{"size-4 | item present", t3, 5, &node{data: 5},  &node{data: 4}},
		{"size-4 | item present (2)", t3, 2, &node{data: 2},  nil},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotParent := tt.tree.Find(tt.find)
			if tt.want != nil && tt.wantParent != nil {
				if got.data != tt.want.data {
					t.Errorf("\nNodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.want.data, got.data)
				}
				if gotParent.data !=tt.wantParent.data {
					t.Errorf("\nParent Nodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.wantParent.data, gotParent.data)
				}
			}
				if tt.want != nil && tt.wantParent == nil {
					if got.data != tt.want.data {
						t.Errorf("Nodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.want.data, got.data)
					}
				}
		})
	}
}

func TestDelete(t *testing.T) {
	t0 := bst{}
	t1 := bst{}
	t1.Insert(2)
	t2 := bst{}
	t2.Insert(2)
	t2.Insert(1)
	t3 := bst{}
	t3.Insert(2)
	t3.Insert(1)
	t3.Insert(4)
	t3.Insert(5)

	tests := []struct {
		name string
		tree bst
		find int
		want *node
		wantParent *node
	}{
		{"empty", t0, 3, nil, nil},
		{"size-1", t1, 3, nil, nil},
		{"size-1 | item present", t1, 2, &node{data: 2}, nil},
		{"size-2 | item present", t2, 1, &node{data: 1},  &node{data: 2}},
		{"size-2 | item NOT present", t2, 3, nil, nil},
		{"size-4 | item present", t3, 5, &node{data: 5},  &node{data: 4}},
		{"size-4 | item present (2)", t3, 2, &node{data: 2},  nil},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotParent := tt.tree.Find(tt.find)
			if tt.want != nil && tt.wantParent != nil {
				if got.data != tt.want.data {
					t.Errorf("\nNodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.want.data, got.data)
				}
				if gotParent.data !=tt.wantParent.data {
					t.Errorf("\nParent Nodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.wantParent.data, gotParent.data)
				}
			}
				if tt.want != nil && tt.wantParent == nil {
					if got.data != tt.want.data {
						t.Errorf("Nodes not equal:\n\twant:%#v\n\tgot: %#v\n", tt.want.data, got.data)
					}
				}
		})
	}
}