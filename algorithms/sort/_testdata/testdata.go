package _testdata

var TestData = []struct {
	Name string
	Input []int
	Want []int
} {
		{"empty", []int{}, []int{}},
		{"size-1", []int{1}, []int{1}},
		{"size-2", []int{1, 2}, []int{1, 2}},
		{"size-2 | unsorted", []int{2, 1}, []int{1, 2}},
}