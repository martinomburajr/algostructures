package main

// type Traverser interface {
// 	BFT() []*Node
// 	InorderDFT() []*Node
// 	PostOrderDFT() []*Node
// 	PreOrderDFT() []*Node
// }

// type Storer interface {
// 	Add(node Node)
// 	Get(id string)
// 	Delete(id string)
// 	Edit(id string)
// }

// type DataStructure interface {
// 	Size() int
// }

type TreeNode struct {
	id string
	value int
	left *Node
	right *Node
}

type BinaryTree struct {
	root *TreeNode
}

// BFS (BreadthFirstSearch)
func (b *BinaryTree) BFS() []*TreeNode {
	if b.root == nil {
		panic ("tree root is nil")
	}
	return nil
}


