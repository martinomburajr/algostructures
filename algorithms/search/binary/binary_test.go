package binary

import (
	"testing"
	"github.com/martinomburajr/algostructures/algorithms/search/_searchtestdata"
)

func TestBinarySearchRecursive(t *testing.T) {
	tests := _searchtestdata.SearchTestData
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := BinarySearchRecursive(tt.Input, tt.Item)
			if got != tt.Want {
				t.Errorf("\n\nwant:\t%d\ngot:\t%d", tt.Want, got)
			}
		})
	}
}

func TestBinarySearchImperative(t *testing.T) {
	tests := _searchtestdata.SearchTestData
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := BinarySearchImperative(tt.Input, tt.Item)
			if got != tt.Want {
				t.Errorf("\n\nwant:\t%d\ngot:\t%d", tt.Want, got)
			}
		})
	}
}