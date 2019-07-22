package utils

import (
	"testing"
	"reflect"
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
			if ans := GenerateOrderedAscendingArray(v.args.start, v.args.size); !CheckSliceEquality(t, ans, v.result) {
				t.Errorf("Wanted %d | got %d", v.result, ans)
			}
		})
	}
}


func TestShift(t *testing.T) {
	type args struct {
		arr []int
		start, size int
	}
	tt := []struct{
		name string
		args args
		result []int
	}{
		{"size-1", args{[]int{1},0,1}, slice1},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if ans := ShiftArray(v.args.start, v.args.size); reflect.DeepEquals(v.result, ans) {
				t.Errorf("Wanted %d | got %d", v.result, ans)
			}
		})
	}
}

