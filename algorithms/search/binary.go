package search

// BinarySearch performs a binary search operation to find the item in the arr. arr must be a sorted ascending array, If it fails it returns the value of -1
func BinarySearch(arr []int, item int) int {
	if len(arr) < 3 {
		return SequentialSearch(arr, item)
	}

	start := 0
	end := len(arr)-1
	mid := (start + end) / 2

	for start < mid  && mid < end {
		if arr[start] == item {
			return start;
		}
		if arr[end] == item {
			return end;
		}
		if arr[mid] == item {
			return mid
		}

		if item > mid {
			start = mid+1
			end = end-1
			mid = (start + end) / 2
		}
		if item < mid {
			start = start + 1
			end = mid-1
			mid = (start + end)/2
		}
	}

	return -1
}
