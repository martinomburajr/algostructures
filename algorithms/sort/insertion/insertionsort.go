package insertion

// InsertionSort
func InsertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		j := i
		for j >= 1 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
	}
	return arr
}