package binary

// BinarySearchRecursive uses recursion to find the item
func BinarySearchRecursive(arr []int, item int) int {
	if len(arr) < 1 {
		return -1
	}

	start := 0
	end := len(arr)-1
	mid := (start+end)/2

	return binarySearch(arr, item, start, mid, end)
}

func binarySearch(arr []int, item, start, mid, end int) int {
	if arr[start] == item {
		return start
	}
	if arr[mid] == item {
		return mid
	}
	if arr[end] == item {
		return end
	}

	if start == mid || mid == end {
		return -1
	}

	if arr[mid] > item {
		start = start+1
		end = mid-1
		mid = (start+end)/2

		return binarySearch(arr, item, start, end, mid)
	}
	start = mid+1
	end = end-1
	mid = (start+end)/2
	return binarySearch(arr, item, start, end, mid)
}

// BinarySearchImperative performs a binary search, but uses a for loop instead
// of recursion
func BinarySearchImperative(arr []int, item int) int {
	if len(arr) < 1 {
		return -1
	}

	start := 0
	end := len(arr)-1
	mid := (start+end)/2

	for {
		if arr[start] == item {
			return start
		}
		if arr[mid] == item {
			return mid
		}
		if arr[end] == item {
			return end
		}
		if start == mid || mid == end {
			return -1
		}

		if arr[mid] > item {
			start = start+1
			end = mid-1
			mid = (start+end)/2
		} else {
			start = mid+1
			end = end-1
			mid = (start+end)/2
		}
	}
}