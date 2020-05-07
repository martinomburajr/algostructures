package sequential

import (
	"testing"
	"github.com/martinomburajr/algostructures/algorithms/search/_searchtestdata"
)

func TestSequentialSearch(t *testing.T) {
	tests := _searchtestdata.SearchTestData
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := SequentialSearch(tt.Input, tt.Item)
			if got != tt.Want {
				t.Errorf("\nwant:\t%d\ngot:\t%d", tt.Want, got)
			}
		})
	}
}