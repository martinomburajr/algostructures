package search

import (
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	emptyArr := make([]int, 0)

	arrItemDoestExist := make([]int, 1)
	arrItemDoestExist[0] = 1

	arrItemExists := make([]int, 2)
	arrItemExists[0] = 1
	arrItemExists[1] = 0

	type args struct {
		arr  []int
		item int
	}
	tt := []struct {
		name   string
		args   args
		result int
	}{
		{ "empty-arr", args{emptyArr, 0}, -1, },
		{	"item doesnt exist", args{arrItemDoestExist, 0}, -1, },
		{	"item exists", args{arrItemExists, 0}, 1, },
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if resp := SequentialSearch(v.args.arr, v.args.item); resp != v.result {
				t.Errorf("got %d expected %d", resp, v.result)
			}
		})
	}
}

func BenchmarkSequentialTest(b *testing.B) {
	arr := []int{0,1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		SequentialSearch(arr, 7)
	}
}
