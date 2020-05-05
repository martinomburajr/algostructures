package sort

import (
	"testing"
	"reflect"
)


func TestComparer(t *testing.T) {
	tests := []struct {
		name string
		inputA []int
		inputB []int
		want []int
	} {
		{"empty", []int{},[]int{}, []int{}},
		{"size-1", []int{1}, []int{0}, []int{0, 1}},
		{"size-2", []int{1, 0}, []int{2}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := comparer(tt.inputA, tt.inputB)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s\nwant: %#v\ngot:%#v", tt.name, tt.want, got)
			}
		})
	}
}


func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		input []int
		want []int
	} {
		{"empty", []int{}, []int{}},
		{"size-1", []int{1}, []int{1}},
		{"size-2", []int{1, 0}, []int{0, 1}},
		{"size-3", []int{1, 0, 2}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s\nwant: %#v\ngot:%#v", tt.name, tt.want, got)
			}
		})
	}
}

