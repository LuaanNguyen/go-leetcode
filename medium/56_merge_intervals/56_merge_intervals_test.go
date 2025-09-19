package merge_intervals

import (
	"reflect"
	"testing"
)


func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		intervals [][]int
		expectedResult [][]int  
	} {
		{
			name: "Test Case 1",
			intervals: [][]int{{1,3},{1,5},{6,7}},
			expectedResult: [][]int{{1,5},{6,7}},
		},
		{
			name: "Test Case 2",
			intervals: [][]int{{1,3},{2,6},{8,10},{15,18}},
			expectedResult: [][]int{{1,6},{8,10},{15,18}},
		},
		{
			name: "Test Case 3",
			intervals: [][]int{{1,4},{4,5}},
			expectedResult: [][]int{{1,5}},
		},
		{
			name: "Test Case 4 (edge case)",
			intervals: [][]int{},
			expectedResult: [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := merge(test.intervals)
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Merge(%v) = %v; want = %v", test.intervals, got, test.expectedResult)
			}
		})

	}
}