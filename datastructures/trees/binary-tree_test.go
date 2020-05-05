package trees

import (
	"testing"
)

var (
	tree0 = BinaryTree{}
	tree1 = BinaryTree{root: &BinaryTreeNode{value: 0}}
	tree2 = BinaryTree{
		root: &BinaryTreeNode{value: 0,
			left: &BinaryTreeNode{value: 1}}}
	tree3 = BinaryTree{
		root: &BinaryTreeNode{value: 0,
			left: &BinaryTreeNode{value: 1,
				left: &BinaryTreeNode{value: 2}}}}
	tree4 = BinaryTree{
		root: &BinaryTreeNode{value: 0,
			left: &BinaryTreeNode{value: 1,
				left: &BinaryTreeNode{value: 2},
				right: &BinaryTreeNode{value: 3}}}}
	tree5 = BinaryTree{
		root: &BinaryTreeNode{value: 0,
			left: &BinaryTreeNode{value: 1,
				left: &BinaryTreeNode{value: 2},
				right: &BinaryTreeNode{value: 3}},
			right: &BinaryTreeNode{value: 4}}}
	tree6 = BinaryTree{
		root: &BinaryTreeNode{value: 0,
			left: &BinaryTreeNode{value: 1,
				left: &BinaryTreeNode{value: 2},
				right: &BinaryTreeNode{value: 3}},
			right: &BinaryTreeNode{value: 4,
				right: &BinaryTreeNode{value: 5}}}}
)

func TestInorderDFS(t *testing.T) {
	tests := []struct {
		name string
		binaryTree BinaryTree
		want []*BinaryTreeNode
	}{
		{"empty", tree0, nil},
		{"tree1", tree1, []*BinaryTreeNode{&BinaryTreeNode{value: 0}}},
		{"tree2", tree2, []*BinaryTreeNode{&BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 0}}},
		{"tree3", tree3, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 0}}},
		{"tree4", tree4, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}}},
		{"tree5", tree5, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 4}}},
		{"tree6", tree6, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 4}, &BinaryTreeNode{value: 5}, }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.binaryTree.InorderDFS()
			if len(got) != len(tt.want) {
				nodePrint(tt.want)
				nodePrint(got)
				t.Errorf("Not the same length")
			}
			for i := range got {
				if got[i].value != tt.want[i].value {
					t.Errorf("\n\nwant: %#v\ngot: %#v at index %d", tt.want[i].value, got[i].value, i)
				}
			}
		})
	}
}

func TestPreorderDFS(t *testing.T) {
	tests := []struct {
		name string
		binaryTree BinaryTree
		want []*BinaryTreeNode
	}{
		// {"empty", tree0, nil},
		// {"tree1", tree1, []*BinaryTreeNode{&BinaryTreeNode{value: 0}}},
		// {"tree2", tree2, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}}},
		// {"tree3", tree3, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 2}}},
		{"tree4", tree4, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 3}}},
		// {"tree4", tree4, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}}},
		// {"tree5", tree5, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 4}}},
		// {"tree6", tree6, []*BinaryTreeNode{&BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 4}, &BinaryTreeNode{value: 5}, }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.binaryTree.PreorderDFS()
			if len(got) != len(tt.want) {
				nodePrint(tt.want)
				nodePrint(got)
				t.Errorf("Not the same length")
			}
			for i := range got {
				if got[i].value != tt.want[i].value {
					t.Errorf("\n\nwant: %#v\ngot: %#v at index %d", tt.want[i].value, got[i].value, i)
				}
			}
		})
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		name string
		binaryTree BinaryTree
		want []*BinaryTreeNode
	}{
		{"empty", tree0, nil},
		{"tree1", tree1, []*BinaryTreeNode{&BinaryTreeNode{value: 0}}},
		{"tree2", tree2, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}}},
		{"tree3", tree3, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 2}}},
		{"tree4", tree4, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 3}}},
		{"tree5", tree5, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 4}, &BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 3} }},
		{"tree6", tree6, []*BinaryTreeNode{&BinaryTreeNode{value: 0}, &BinaryTreeNode{value: 1}, &BinaryTreeNode{value: 4}, &BinaryTreeNode{value: 2}, &BinaryTreeNode{value: 3}, &BinaryTreeNode{value: 5} }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.binaryTree.BFS()
			if len(got) != len(tt.want) {
				nodePrint(tt.want)
				nodePrint(got)
				t.Errorf("Not the same length")
			}
			for i := range got {
				if got[i].value != tt.want[i].value {
					t.Errorf("\n\nwant: %#v\ngot: %#v at index %d", tt.want[i].value, got[i].value, i)
				}
			}
		})
	}
}