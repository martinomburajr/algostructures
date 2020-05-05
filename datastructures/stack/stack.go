package stack

// Node represents the data stored in a stack
type Node struct {
	value int
	next *Node
}

// Stack represents a LIFO datastructure that resembles a list
type Stack struct {
	head *Node
}

// Size returns the size of the Stack
func (s *Stack) Size() int {
	i := 0
	curr := s.head
	for {
		if curr == nil {
			return i
		}
		i++
		curr = curr.next
	}
}

// Pop removes the first item from the stack
func (s *Stack) Pop() *Node {
	if s.head == nil {
		return nil
	}
	curr := s.head
	s.head = s.head.next
	return curr
}

// Push adds an item to the front of the stack
func (s *Stack) Push(data int) {
	node := &Node{value: data, next: s.head}
	s.head = node
}

// Peek retrieves the first item in the stack
func (s *Stack) Peek() *Node {
	return s.head
}

// ClearStack unravels the stack and returns a slice of its elements in LIFO order
func (s *Stack) ClearStack() []*Node {
	size := s.Size()
	nodes := make([]*Node, size)

	i := 0
	curr := s.head
	for {
		if curr == nil {
			return nodes
		}
		nodes[(size-1)-i] = curr

		i++
		curr = curr.next
	}
}