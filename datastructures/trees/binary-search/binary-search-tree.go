package binary_search

import (
	"fmt"
)

type IBST interface {
	Insert(data int)
	Delete(data int)
	Find(data int) (currNode *node, parentNode *node)
}

type ITraverser interface {
	Size()
	InorderDFS()
	PreorderDFS()
	PostorderDFS()
	BFS()
}

type node struct {
	data int
	left *node
	right *node
}

type bst struct {
	root *node
}

// Insert adds an item to the Binary Search Tree
func (b *bst) Insert(data int) {
	if b.root == nil {
		node := &node{data: data}
		b.root = node
		return
	}

	insert(b.root, nil, data)
}

func (b *bst) Find(data int) (currNode *node, parentNode *node) {
	return find(b.root, nil, data)
}

func find(currNode *node, parentNode *node, data int) (node *node, parent *node) {
	if currNode == nil {
		return
	}

	if currNode.data > data {
		node, parent = find(currNode.left, currNode, data)
	} else if currNode.data < data {
		node, parent = find(currNode.right, currNode, data)
	} else {
		node = currNode
		parent = parentNode
		return node, parent
	}
	return node, parent
}

// Delete performs a delete operation on any given data.
func (b *bst) Delete(data int) {
	node, parent := b.Find(data)

	// leaf node
	if node.left == nil && node.right == nil {
		if parent.left.data == data {
			parent.left = nil
		}
		if parent.right.data == data {
			parent.right = nil
		}
		return
	}

	// one child
	if node.left != nil && node.right == nil {
		parent.left = node.left
	}
	if node.right != nil && node.left == nil {
		parent.right = node.right
	}

	// two children
	if node.left != nil && node.right != nil {
		if parent.data >= data {
			child := node.left
			child.right = node.right
			parent.left = child
		} else {
			child := node.left
			child.right = node.right
			parent.right = child
		}
	}
}

func insert(currNode *node, parent *node, data int) {
	if currNode != nil {
		if currNode.data >= data {
			insert(currNode.left, currNode, data)
		} else {
			insert(currNode.right, currNode, data)
		}
	} else {
		newNode := &node{data: data}
		if parent.data >= data {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}
}

func (b *bst) InorderDFS() []*node {
	if b.root == nil {
		return nil
	}

	traversalList := make([]*node, 0)
	inorderDFS(b.root, &traversalList)

	nodePrint(traversalList)
	return traversalList
}

func inorderDFS(node *node, traversalList *[]*node) {
	if node == nil {
		return
	}
	inorderDFS(node.left, traversalList)
	*traversalList = append(*traversalList, node)
	inorderDFS(node.right, traversalList)
}

func nodePrint(arr []*node) {
	for i := range arr {
		fmt.Printf("%#v, ", arr[i].data)
	}
	fmt.Println()
	fmt.Println()
}