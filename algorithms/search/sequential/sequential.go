package sequential

// SequentialSearch perfroms a linear search against the input array for the
// item. It returns a positive index (including 0), otherwise a -1 if not found.
func SequentialSearch(arr []int, item int) int {
	if len(arr) < 1 {
		return -1
	}

	foundIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			foundIndex = i
			break
		}
	}

	return foundIndex
}