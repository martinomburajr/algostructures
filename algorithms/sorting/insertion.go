package sorting

import (
	"go/format"
)

// InsertionSort is a comparison-based algorithm that builds a final sorted array one element at a time. It iterates through an input array and removes one element per iteration, finds the place the element belongs in the array, and then places it there.
func InsertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	tracer := 0
	for {
		arr[tracer] < 1 {}
	}
}

// ShiftArray will take an item from a given index and move items between the from and to index, and place the item at the to index. Example arr=[3,4,6,7,5] from =4 to = 2. Output is [3,4,5,6,7]. This method does have sideffects.
func ShiftArray(arr *[]int, from, to int) {
	tempArrStart := arr[0:to]
	tempArrNext := arr[to:from]
	finArr := append(tempArrStart, arr[from], tempArrNext...)
}