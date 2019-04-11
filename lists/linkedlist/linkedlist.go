package linkedlist

import "fmt"

//LinkedList is a type of list that comprises of nodes that have pointers to the next Node.
// This version of a linked list is not doubly-linked, so each node can have at most 1 pointers to it.
type LinkedList struct {
	count int64
	head *node
}

func (l *LinkedList) Get(index int64) (interface{}, error) {
	if index == 0 && l.count == 0 {
		return nil, fmt.Errorf("list is empty / uninitialized")
	}
	if index >= l.count {
		return nil, fmt.Errorf("index out of bounds (upper)")
	}

	if index == 0 && l.count > 0 {
		return l.head.v, nil
	}
	if index > 0 {
		var i int64
		temp := l.head.next
		for i = 1; i <= index; i++  {
			if i == index {
				return temp.v, nil
			}
			temp = temp.next
		}
	}
	 return nil, fmt.Errorf("illegal negative index")
}

func (l *LinkedList) Add(item interface{}) {
	if l.count == 0 {
		l.head = &node{item, nil}
		l.count++
	} else if  l.count == 1 {
		l.head.next = &node{item, nil}
		l.count++
	}else if l.count > 1 {
		//
		l.count++
	}
}

func (l *LinkedList) Size() int64 {
	return l.count
}