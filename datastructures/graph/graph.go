package graph

import (
	"fmt"
)
// Node represents the data stored in a stack
type Node struct {
	value *GraphNode
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
func (s *Stack) Push(data *GraphNode) {
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

type GraphNode struct {
	value int
	nodes []*GraphNode
}

type Graph struct {
	nodes []*GraphNode
}

func (g *Graph) BFS () []*GraphNode {
	return nil
}

func (g *Graph) DFS () []*GraphNode {
	if g.nodes == nil {
		return nil
	}

	visitedMap := make(map[*GraphNode]bool)
	visitedStack := &Stack{head: &Node{value: g.nodes[0], next: nil}}
	traversedNodes := make([]*GraphNode, 0)

	dfs(g.nodes[0], visitedMap, visitedStack, &traversedNodes)

	return traversedNodes
}

func dfs(node *GraphNode, visitedMap map[*GraphNode]bool, visitedStack *Stack, traversedNodes *[]*GraphNode) {
	if node.nodes != nil {
		for i := range node.nodes {
			innerNode := node.nodes[i]
			if visitedMap[innerNode] {
				continue
			} else {
				visitedStack.Push(innerNode)
				visitedMap[innerNode] = true
			}
		}
	}

	stackNodes := visitedStack.ClearStack()
	fmt.Printf("PRINT-ALL: %#v\n", stackNodes)
	visitedStack = &Stack{}
	for i := range stackNodes {
		*traversedNodes = append(*traversedNodes, stackNodes[i].value)
		fmt.Printf("PRINT: %#v\n", stackNodes[i].value)
	}
	return
}