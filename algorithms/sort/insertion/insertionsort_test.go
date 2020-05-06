package insertion

import (
	"testing"
	"reflect"
	"github.com/martinomburajr/algostructures/algorithms/sort/_testdata"
)

func TestInsertionSort(t *testing.T) {
	tests := _testdata.TestData
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := InsertionSort(tt.Input)
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("\n\nwant:\t%#v\ngot:\t%#v", tt.Want, got)
			}
		})
	}
}