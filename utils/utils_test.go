package utils

import (
	"testing"
)

func TestGenerateOrderedAscendingArray(t *testing.T) {
	// emptySlice := make([]int, 0, 0)
	slice1 := make([]int, 1,1)
	type args struct {
		start, size int
	}
	tt := []struct{
		name string
		args args
		result []int
	}{
		// {"size-0", args{0,0}, emptySlice},
		{"size-1", args{0,1}, slice1},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if ans := GenerateOrderedAscendingArray(v.args.start, v.args.size); checkSliceEquality(t, ans, v.result) {
				t.Errorf("Wanted %d | got %d", v.result, ans)
			}
		})
	}
}

func checkSliceEquality(t *testing.T, sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		t.Errorf("\nsliceA and sliceB are not the same length | %#v - %#v \b", sliceA, sliceB)
		return false
	}
	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			t.Errorf("\ncontain different items at index %d | %#v - %#v \b", i, sliceA, sliceB)
			return false
		}
	}
	return true
}