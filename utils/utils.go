package utils

import( 
	"fmt"
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
