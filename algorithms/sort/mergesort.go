package sort

// MergeSort is a divide and conquer algorithm that sorts an array recursively.
// O(nlog(n)) - Omega(nlog(n)) - Theta(nlog(n))
func MergeSort(arr []int) []int {
	//1. Divide & Conquer
	if len(arr) > 1 {
		leftArr := arr[:len(arr)/2]
		rightArr := arr[len(arr)/2:]

		leftSide := MergeSort(leftArr)
		rightSide := MergeSort(rightArr)

		outputArr := []int{}
		if len(leftSide) > len(rightSide) {
			for i := 0; i < len(leftSide); i++ {
				for j := 0; j < len(rightSide); j++ {
					if leftSide[i] > rightSide[j] {
						outputArr = append(outputArr, rightSide[j])
						outputArr = append(outputArr, leftSide[i])
					} else {
						outputArr = append(outputArr, leftSide[i])
						outputArr = append(outputArr, rightSide[j])
					}
				}
			}
		} else {
			for j := 0; j < len(rightSide); j++ {
				for i := 0; i < len(leftSide); i++ {
					if leftSide[i] > rightSide[j] {
						outputArr = append(outputArr, rightSide[j])
						outputArr = append(outputArr, leftSide[i])
					} else {
						outputArr = append(outputArr, leftSide[i])
						outputArr = append(outputArr, rightSide[j])
					}
				}
			}
		}
		return outputArr
	}

	return arr
}

func comparer(a, b []int) []int {
	finArr := []int{}
	if len(a) >= len(b) {
		var i, j = 0, 0

		for i < len(a) {
			for j < len(b) {
				if a[i] > b[j] {
					finArr = append(finArr, b[j])
					if i == len(a) - 1 {
						finArr = append(finArr, a[i])
						break
					}
					j=j+1
				} else {
					finArr = append(finArr, a[i])
					if j == len(b) - 1 {
						finArr = append(finArr, b[j])
						break
					}
					i=i+1
					break
				}
			}
			i++
		}
	}

	return finArr
}