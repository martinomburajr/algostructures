package heap

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name string
		heap MinHeap
		indexA int
		indexB int
		want []int
		shouldPanic bool
	} {
		{"add to empty", MinHeap{heap: []int{}}, 0, 1, []int{1}, true},
		{"panic - array small", MinHeap{heap: []int{1}}, 0, 1, []int{1}, true},
		{"array ok | size 2", MinHeap{heap: []int{1, 2}}, 0, 1, []int{2, 1}, false},
		{"array ok | size 4", MinHeap{heap: []int{1, 2,3,4}}, 0, 3, []int{4, 2, 3, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := swap(tt.heap.heap, tt.indexA, tt.indexB)

			if len(tt.want) != len(got) {
				t.Errorf("heap lengths not the same:\nwant:\t%d\ngot:\t%d", len(tt.want), len(got))
			}
			for i := range tt.heap.heap {
				if tt.want[i] != tt.heap.heap[i] {
					t.Errorf("heap item not the same:\nwant:\t%#v\ngot:\t%#v", tt.want, got)
				}
			}
		})
	}
}

func TestHeapAdd(t *testing.T) {
	tests := []struct {
		name string
		heap MinHeap
		item int
		want []int
	} {
		{"add to empty", MinHeap{heap: []int{}}, 1, []int{1}},
		{"add to size-1", MinHeap{heap: []int{0}}, 1, []int{0,1}},
		{"add to size-1 (smaller)", MinHeap{heap: []int{0}}, -1, []int{-1, 0}},
		{"add to size-2", MinHeap{heap: []int{0, 1}}, 2, []int{0,1,2}},
		{"add to size-2", MinHeap{heap: []int{0, 1}}, -1, []int{-1,1, 0}},
		{"add to size-3", MinHeap{heap: []int{0, 1, 3}}, -1, []int{-1,0, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Add(tt.item)

			if len(tt.want) != len(tt.heap.heap) {
				t.Errorf("heap lengths not the same:\nwant:\t%d\ngot:\t%d", len(tt.want), len(tt.heap.heap))
			}
			for i := range tt.heap.heap {
				if tt.want[i] != tt.heap.heap[i] {
					t.Errorf("heap item not the same:\nwant:\t%#v\ngot:\t%#v", tt.want, tt.heap.heap)
				}
			}
		})
	}
}

func TestHeapPop(t *testing.T) {
	tests := []struct {
		name string
		heap MinHeap
		wantPop int
		want []int
		shouldPanic bool
	} {
		{"pop empty", MinHeap{heap: []int{}}, 1, []int{1}, true},
		{"pop size-1", MinHeap{heap: []int{0}}, 0, []int{}, false},
		{"pop size-2", MinHeap{heap: []int{0, 1}}, 0, []int{1}, false},
		// {"add to size-2", MinHeap{heap: []int{0, 1}}, -1, []int{-1,1, 0}},
		// {"add to size-3", MinHeap{heap: []int{0, 1, 3}}, -1, []int{-1,0, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			got := tt.heap.Pop()

			if tt.wantPop != got {
				t.Errorf("pop items:\nwant:\t%d\ngot:\t%d", tt.wantPop, got)
			}
			if len(tt.want) != len(tt.heap.heap) {
				t.Errorf("heap lengths not the same:\nwant:\t%d\ngot:\t%d", len(tt.want), len(tt.heap.heap))
			}
			for i := range tt.heap.heap {
				if tt.want[i] != tt.heap.heap[i] {
					t.Errorf("heap item not the same:\nwant:\t%#v\ngot:\t%#v", tt.want, tt.heap.heap)
				}
			}
		})
	}
}