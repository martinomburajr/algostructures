package sequential

import "testing"

func TestSequentialTable(t *testing.T) {
	tests  := []struct {
		expected int
		input int
	} {
		{ 3, SequentialSearch([]int{1,4,6,77,88,99}, 77) },
	}

	for _, test := range tests {
		if test.expected != test.input {
			t.Error("Test failed!")
		}
	}
}

