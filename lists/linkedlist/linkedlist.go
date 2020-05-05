package linkedlist

import "fmt"

//LinkedList is a type of list that comprises of nodes that have pointers to the next Node.
// This version of a linked list is not doubly-linked, so each node can have at most 1 pointers to it.
type LinkedList struct {
	count int64
	head *node
}

//Get retrieves an item from a given index within the list
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

//Add will add an item to the end of the list
func (l *LinkedList) Add(item interface{}) {
	if l.count == 0 {
		l.head = &node{item, nil}
		l.count++
	} else if  l.count == 1 {
		l.head.next = &node{item, nil}
		l.count++
	}else if l.count > 1 {
		var i int64
		temp := l.head.next
		for i = 1; i <= l.count; i++  {
			if i == l.count-1 {
				temp.next = &node{item, nil}
				return
			}
			temp = temp.next
		}
		l.count++
	}
}

//Remove will remove an item from the specified index and return the count of the list.
func (l *LinkedList) Remove(index int64) (int64, error) {
	if l.count == 0 {
		return -1, fmt.Errorf("cannot remove from empty list")
	}

	if index > l.count {
		return -1, fmt.Errorf("index out of bounds (upper)")
	}else if index < 0 {
		return -1, fmt.Errorf("index out of bounds (lower)")
	}

	var i int64
	temp := l.head
	for i = 0; i <= index ; i++  {
		if i == index-1 {
			temp.next = temp.next.next
			temp.next.next = nil
			l.count--
			return l.count, nil
		}
		temp = temp.next
	}

	return 0, nil
}

//Size returns the current size of the list.
func (l *LinkedList) Size() int64 {
	return l.count
}

//ToSlice returns a slice version of the linkedlist. If the linkedlist has size of 0 it returns nil
func (l *LinkedList) ToSlice() []interface{} {
	if l.count <= 0 {
		return nil
	}

	s := make([]interface{}, l.count)
	var i int64
	temp := l.head
	for i = 0; i < l.count; i++ {
		s = append(s, temp.v)
		if i != l.count - 1 {
			temp = temp.next
		}
	}
	return s
}