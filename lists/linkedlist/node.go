package linkedlist

//node forms part of a linked list and carries the relevant information of it
type node struct {
	v interface{}
	next *node
}