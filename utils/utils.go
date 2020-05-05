package utils

import (
	"fmt"
	"testing"
)

// GenerateOrderedAscendingArray creates an ascending array of ordered integers from a given start number + the size.
func GenerateOrderedAscendingArray(start, size int) []int {
	arr := make([]int, size, size)
	for i := start; i < size; i++ {
		arr[i] = i
	}

	fmt.Printf("%#v", arr)
	return arr
}

// CheckSliceEquality uses basic checks to check slice equality
func CheckSliceEquality(t *testing.T, expected, got []int) bool {
	if len(expected) != len(got) {
		t.Errorf("\nexpected %#v | got %#v - not the same length", expected, got)
		return false
	}
	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("\nexpected %#v | got %#v - contain different items at index %d", i, expected, got)
			return false
		}
	}
	return true
}

// ShiftArray will take an item from a given index and move items between the from and to index, and place the item at the to index. Example 
// arr=[3,4,6,7,5,8,9] from=4 to=2. Output is [3,4,5,6,7,8,9].
// arr=[3,4,6,7,5] from=4 to=2. Output is [3,4,5,6,7].
// arr=[2,6,4,5] from=1 to=3. Output is [2,4,5,6]
// arr=[2,6,4] from=1 to=2. Output is [2,4,6]
// arr=[2,1] from=1 to=0. Output is [1,2]
// arr=[1] from=0 to=0. Output is [1,2]
// It can shift both forward and backwards. Invalid from and to parameters e.g. from and to are out of bounds will return nil
func ShiftArray(arr []int, from, to int) []int {
	// check indices are not out of bounds
	if from < 0 || from > len(arr) {
		return nil
	}
	if to < 0 || to > len(arr) {
		return nil
	}
	// check cases where its okay to return the input array
	if len(arr) < 2 || from == to {
		return arr
	}
	
	if from > to {
		tempArrStart := arr[0:to]
		tempArrNext := arr[to:from]
		finArr := append(tempArrStart, arr[from])
		finArr = append(finArr,  tempArrNext...)
		return finArr
	}
	
	return nil
}
