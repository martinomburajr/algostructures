package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		item int
	}
	tt := []struct{
		name string
		args args
		result int
	}{
		{"empty array", args{[]int{}, 1}, -1},
		{"1 item array", args{[]int{1}, 1}, 0},
		{"1 item array without", args{[]int{2}, 1}, -1},
		{"2 item array", args{[]int{1,2}, 2}, 1},
		{"2 item array without", args{[]int{1,2}, 0}, -1},
		{"3 item array", args{[]int{1,2,3},3}, 2},
		{"3 item array without", args{[]int{1,2,3}, 4}, -1},
		{"4 item array", args{[]int{1,2,3,4}, 1}, 0},
		{"4 item array without", args{[]int{1,2,3,4}, 5}, -1},
		{"multi item array with", args{[]int{0,1,2,3,4,5,6,7,8,9,10}, 9}, 9},
		{"multi item array without", args{[]int{0,1,2,3,4,5,6,7,8,9,10}, 11}, -1},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if ans := BinarySearch(v.args.arr, v.args.item); ans != v.result {
				t.Errorf("Wanted %d | got %d", v.result, ans)
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{0,1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 7)
	}
}