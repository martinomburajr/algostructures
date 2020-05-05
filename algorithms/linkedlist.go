package main

import (

)

type ListTraverser interface {
	Traverse(index int) *Node
}

type Storer interface {
	Add(node Node)
	Get(id string)
	Delete(id string)
	Edit(id string)
}

type DataStructure interface {
	Size() int
}

type Node struct {
	id string
	value int
	next *Node
	right *Node
}

type LinkedList struct {
	root *Node
}

// Get returns a node at a given index
func (l *LinkedList) Get(index int) *Node {
	if l.root == nil {
		panic("root is nil")
	}

	if index == 0 {
		return l.root
	}

	next := l.root.next
	if next == nil {
		panic("index out of bounds")
	}

	for i := 1; i < index; i++ {
		next = next.next
		if next == nil {
			panic("index out of bounds")
		}
	}

	return next
}

// Size refers to the number items in the linked list
func (l *LinkedList) Size() int {
	if l.root == nil {
		panic("root is nil")
	}

	if l.root.next == nil {
		return 0
	}

	next := l.root.next
	counter := 1
	for {
		if next == nil {
			return counter
		}
		next = next.next
		counter++
	}
}
//Add adds to the end of the linked list
func (l *LinkedList) Add(node Node) {
	if l.root == nil {
		l.root = &node
		return
	}

	next := l.root.next
	prev := l.root
	for {
		if next == nil {
			prev.next = &node
			return
		}
		prev = next
		next = next.next
	}
}

// func (l *LinkedList) Delete(index int) {
// 	if l.root == nil {
// 		panic("root is nil")
// 	}
// 	if index == 0 {
// 		l.root = l.root.next
// 		return
// 	}

// 	prev := l.root
// 	next := l.root.next

// 	for i := 1; i < index; i++{
// 		if next.next == nil {
// 			panic ("index out of bounds")
// 		}
// 		if i == index-1 {
// 			prev.next = nil
// 			prev.next.next
// 		}
// 		prev = next
// 		next = next.next
// 	}
// }