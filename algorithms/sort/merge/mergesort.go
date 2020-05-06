package mergesort

// MergeSort is a divide and conquer algorithm that sorts an array recursively.
// O(nlog(n)) - Omega(nlog(n)) - Theta(nlog(n))
func MergeSort(arr []int) []int {
	//1. Divide & Conquer
	if len(arr) > 1 {
		leftArr := arr[:len(arr)/2]
		rightArr := arr[len(arr)/2:]

		leftSide := MergeSort(leftArr)
		rightSide := MergeSort(rightArr)

		outputArr := comparer(leftSide, rightSide)
		return outputArr
	}

	return arr
}

func comparer(left, right []int) []int {
	newArr := make([]int, 0)

	var i, j = 0, 0
	for {
		if i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				newArr = append(newArr, left[i])
				i++
			} else {
				newArr = append(newArr, right[j])
				j++
			}
		}

		if i >= len(left) {
			newArr = append(newArr, right[j:]...)
			break
		}
		if j >= len(right) {
			newArr = append(newArr, left[i:]...)
			break
		}
	}

	return newArr
}