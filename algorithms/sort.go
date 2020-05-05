package main

import (
	"math/rand"
)

//
func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			next := arr[i+1]
			curr := arr[i]

			arr[i+1] = curr
			arr[i] = next
			j := 1
			for {
				descendingIndex := (i-j)
				if descendingIndex > -1 && next < arr[descendingIndex] {
					swap(arr, (descendingIndex), descendingIndex+1)
					next = arr[descendingIndex]
				} else {
					break
				}
				j++
			}
		}
	}

	return arr
}

func swap(arr []int, i, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}

// BubbleSort sorts an array by iteratively comparing adjacent items in the
// array and moving the larger item further down the array to create an
// ascending list of ints.
func BubbleSort(arr []int)  []int{
	if len(arr) <= 1 {
		return arr
	}
	for {
		hasSwapped := false
		i := 0
		curr := arr[i]
		for j := 1; j < len(arr); j++ {
			if curr > arr[j] {
				temp := arr[j]
				arr[j] = curr
				arr[i] = temp

				hasSwapped = true
			} else {
				curr = arr[i+1]
			}
			i++
		}
		if !hasSwapped {
			break
		}
	}
	return arr
}

// MergeSort uses recursion to split the input array into smaller arrays until
// the smalles array possible is of length 2 (for even sized input array) or 1.
// It then compares adjacent elements and uses the stack to percolate up sorted
// arrays in the order in which they were broken down.
func MergeSort(arr []int) []int {
	if len(arr) <= 2 {
		if len(arr) == 2 {
			if arr[0] > arr[1] {
				temp := arr[1]
				arr[1] = arr[0]
				arr[0] = temp
			}
		}
		return arr
	} else {
		leftArr := MergeSort(arr[:len(arr)/2])
		rightArr := MergeSort(arr[len(arr)/2:len(arr)])
		arr = append(leftArr, rightArr ...)
		return arr
	}
}

// QuickSort performs sorting by randomly selecting a pivot point and using it
// to order items smaller than it on the left, and items larger than it, on the right
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) > 1 {
		randIndex := rand.Intn(len(arr))
		rightArr := make([]int, 0)
		leftArr := make([]int, 0)
		pivot := arr[randIndex]
		for i := 0; i < len(arr); i++ {
			if i == randIndex {
				continue
			}
			if arr[i] > arr[randIndex] {
				rightArr = append(rightArr, arr[i])
			} else {
				leftArr = append(leftArr, arr[i])
			}
		}
		leftSide := QuickSort(leftArr)
		leftSide = append(leftSide, pivot)
		rightSide := QuickSort(rightArr)
		// rightSide = append(rightSide, pivot)

		return append(leftSide, rightSide ...)
	} else {
		return arr
	}
}

// func main() {
// 	x := genRandomArray(1000, 10)
// 	// a := []int{1,2,3}
// 	// b := []int{4,5,6}
// 	// c := append(a, b ...)
// 	fmt.Printf("%#v", x)
// }

func genRandomArray(size int, upperRange int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(upperRange)
	}
	return arr
}