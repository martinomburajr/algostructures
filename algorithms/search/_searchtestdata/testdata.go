package _searchtestdata

var SearchTestData = []struct {
	Name string
	Input []int
	Item int
	Want int
} {
	{"empty", []int{}, 0, -1},
	{"size-1 | no item", []int{1}, 0, -1},
	{"size-1 | found item", []int{1}, 1, 0},
	{"size-2 | no item", []int{1,2}, 3, -1},
	{"size-2 | found item", []int{1,2}, 2, 1},
}