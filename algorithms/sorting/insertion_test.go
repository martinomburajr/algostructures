package sorting

import (
	"testing"
	"github.com/martinomburajr/algostructures/utils"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tt := []struct{
		name string
		args args
		result []int
	}{
		{"empty array", args{[]int{}}, []int{}},
		{"array-size-1", args{[]int{1}}, []int{1}},
		{"array-size-2", args{[]int{1,2}}, []int{1,2}},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T){
			if response := InsertionSort(v.args.arr); !utils.CheckSliceEquality(t, v.result, response) {
				t.Errorf("slices are not equal")
			}
		})
	}


}