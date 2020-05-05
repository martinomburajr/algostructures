package trees

import (
	"fmt"
)

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

// InorderDFS traverses a tree using depth first search, in an inorder manner.
func (b *BinaryTree) InorderDFS() []*BinaryTreeNode {
	if b.root == nil {
		return nil
	}
	traversedNodes := make([]*BinaryTreeNode, 0)
	inorderDFS(b.root, &traversedNodes)

	nodePrint(traversedNodes)
	return traversedNodes
}

func inorderDFS(node *BinaryTreeNode, traversedNodes *[]*BinaryTreeNode) {
	if node == nil {
		return
	}
	inorderDFS(node.left, traversedNodes)
	*traversedNodes = append(*traversedNodes, node)
	inorderDFS(node.right, traversedNodes)
}

func nodePrint(arr []*BinaryTreeNode) {
	for i := range arr {
		fmt.Printf("%#v, ", arr[i].value)
	}
}

func (b *BinaryTree) PreorderDFS() []*BinaryTreeNode {
	if b.root == nil {
		return nil
	}
	traversedNodes := make([]*BinaryTreeNode, 0)
	preorderDFS(b.root, &traversedNodes)

	nodePrint(traversedNodes)
	return traversedNodes
}

func preorderDFS(node *BinaryTreeNode, traversedNodes *[]*BinaryTreeNode) {
	if node == nil {
		return
	}
	*traversedNodes = append(*traversedNodes, node)
	preorderDFS(node.left, traversedNodes)
	preorderDFS(node.right, traversedNodes)
}

func (b *BinaryTree) PostorderDFS() []*BinaryTreeNode {
	if b.root == nil {
		return nil
	}
	traversedNodes := make([]*BinaryTreeNode, 0)
	postorderDFS(b.root, &traversedNodes)

	nodePrint(traversedNodes)
	return traversedNodes
}

func postorderDFS(node *BinaryTreeNode, traversedNodes *[]*BinaryTreeNode) {
	if node == nil {
		return
	}
	postorderDFS(node.left, traversedNodes)
	*traversedNodes = append(*traversedNodes, node)
	postorderDFS(node.right, traversedNodes)
}

// BFS performs a breadth first search on the tree
func (b *BinaryTree) BFS() []*BinaryTreeNode {
	if b.root == nil {
		return nil
	}

	queue := make([]*BinaryTreeNode, 1)

	i := 0
	queue[i] = b.root
	for {
		node := queue[i]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}

		if i == len(queue)-1 {
			return queue
		}
		i++
	}
}
