package search

// SequentialSearch iterates through a given array {arr} to find the item of type int
func SequentialSearch(arr []int, item int) (index int) {
	for i := range arr {
		if arr[i] == item {
			return i
		}
	}
	return -1
}