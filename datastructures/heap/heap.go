package heap

import (
)

// Heap is an interface containing common heap operations.
type Heap interface {
	Add()
	Delete()
	Pop()
	Peek()
	percolateUp()
	percolateDown()
}
// MinHeap is a tree-like datastructure that orders its nodes such that the root
// node is always the smallest item in the tree, and each parents children are
// greater than it. Heaps can be visualized as trees but are typically
// implemented using arrays/slices
type MinHeap struct {
	heap []int
}

func (h *MinHeap) getParent(childIndex int) int {
	if childIndex < 2 {
		return 0
	}

	return (childIndex - 1)/2
}


func (h *MinHeap) getChild(parentIndex int) (leftChild int, rightChild int) {
	if parentIndex < 0 {
		return 0, 0
	}
	if len(h.heap) < 3 {
		return 1, -1
	}
	leftChild = (2 * parentIndex) + 1
	rightChild = (2 * parentIndex) + 2

	if leftChild < len(h.heap) && rightChild < len(h.heap) {
		return leftChild, rightChild
	}

	return -1, -1
}

// Add adds a data point to the heap
func (h *MinHeap) Add(data int) {
	if len(h.heap) < 1 {
		h.heap = append(h.heap, data)
		return
	}
	h.heap = append(h.heap, data)
	addedItemIndex := len(h.heap) - 1

	for {
		if addedItemIndex < 1 {
			return
		}
		parentIndex := h.getParent(addedItemIndex)
		if h.heap[parentIndex] > h.heap[addedItemIndex] {
			h.heap = swap(h.heap, parentIndex, addedItemIndex)
			addedItemIndex = parentIndex
		} else {
			break
		}
	}
}

// Pop removes the top item from the heap and rearranges the tree.
func (h *MinHeap) Pop() int {
	if len(h.heap) == 0 {
		panic("cannot pop from empty heap")
	}
	if len(h.heap) == 1 {
		originalMin := h.heap[0]
		h.heap = h.heap[:len(h.heap)-1]
		return originalMin
	}

	originalMin := h.heap[0] // store original

	mostRecentChildIndex := len(h.heap) -1
	h.heap[0] = h.heap[mostRecentChildIndex]
	mostRecentChildIndex = 0

	for {
		childAIndex, childBIndex := h.getChild(mostRecentChildIndex)

		if h.heap[mostRecentChildIndex] > h.heap[childAIndex] {
			if childBIndex == -1 {
					h.heap = swap(h.heap, mostRecentChildIndex, childAIndex)
					mostRecentChildIndex = childAIndex
			} else {
				if h.heap[childAIndex] < h.heap[childBIndex] {
					h.heap = swap(h.heap, mostRecentChildIndex, childAIndex)
					mostRecentChildIndex = childAIndex
				} else {
					h.heap = swap(h.heap, mostRecentChildIndex, childBIndex)
					mostRecentChildIndex = childBIndex
				}
			}
		} else {
			break
		}
	}
	// truncate the size of the heap
	h.heap = h.heap[:len(h.heap)-1]
	return originalMin
}

func swap(heap []int, indexA, indexB int) []int {
	temp := heap[indexA]
	heap[indexA] = heap[indexB]
	heap[indexB] = temp

	return heap
}

func (h *MinHeap) Delete(data int) {

}

func (h *MinHeap) Peek() int {
	return 0
}

func (h *MinHeap) percolateUp() int {
	return 0
}

func (h *MinHeap) percolateDown() int {
	return 0
}