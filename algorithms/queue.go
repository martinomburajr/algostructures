package main

import (
	"fmt"
	"strings"
)

type IQueue interface {
	isEmpty() bool
	peek() int
	add(item int) int
	remove() int
	size() int
}

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	Head *QueueNode
}


func (q *Queue) IsEmpty() bool {
	if q.Head == nil {
		return true
	}
	return false
}

// Size calculates the size of a given Queue
func (q *Queue) Size() int {
	counter := 0
	next := q.Head

	for {
		if next == nil {
			return counter
		}
		next = next.next
		counter++
	}

	return counter
}

// func (q *Queue) Add(data int) (count int) {
// 	curr := q.Head
// 	prev := curr
// 	for {
// 		if curr == nil {
// 			prev.next = &QueueNode{data: data, next: nil}
// 			return count
// 		}
// 		prev = curr
// 		curr = curr.next
// 		count++
// 	}
// }

// Add adds an item to the back of the queue and returns the new size of the queue.
func (q *Queue) Add(data int) (count int) {
	newItem := &QueueNode{data: data}

	curr := q.Head
	prev := curr
	for {
		if curr == nil {
			if count == 0 {
				q.Head = newItem
				count++
				return count
			}
			prev.next = newItem
			count++
			return count
		}
		prev = curr
		curr = curr.next
		count++
	}
}

// Remove takes the first item off the queue
func (q *Queue) Remove() *QueueNode {
	if q.Head == nil {
		return nil
	}
	if q.Head.next == nil {
		return q.Head
	}

	itemToRemove := q.Head
	q.Head = q.Head.next

	itemToRemove.next = nil
	return itemToRemove
}

// traverse moves through the queue, returning a slice of sequentially visited
// QueueNodes. This function is used to help with testing other Queue functions
// for structural integrity.
func (q *Queue) traverse() []*QueueNode {
	traversedNodes := make([]*QueueNode, 0)

	curr := q.Head
	for {
		if curr == nil {
			return traversedNodes
		}
		traversedNodes = append(traversedNodes, curr)
		curr = curr.next
	}
}

func (q *Queue) print() string {
	sb := strings.Builder{}
	queueNodes := q.traverse()
	for i, nodes := range queueNodes {
		sb.WriteString(fmt.Sprintf("[%d] -> [%#v]\n", i, nodes))
	}

	return sb.String()
}



func main() {
// 	q := Queue{}
// 	q.Head = &QueueNode{data: 0}
// 	q.Head.next = &QueueNode{data: 1}
// 	q.Head.next.next = &QueueNode{data: 2}

// 	fmt.Printf("Queue: %v\n", q)
// 	fmt.Printf("Queue isEmpty(): %t\n", q.IsEmpty())
// 	fmt.Printf("Size: %d\n", q.Size())

// 	q.Add(3)
// 	count := q.Add(4)
// 	fmt.Printf("Size: %d \t Count: %d\n", q.Size(), count)

	str := "4321"
	fmt.Printf("%t", str[0] > str[3])
}