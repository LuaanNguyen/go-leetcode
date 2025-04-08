package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name 	string 
		nums	[]int 
		expectedResult 	[][]int
	} {
		{
			name: "Test Case 1",
			nums: []int{-1,0,1,2,-1,-4},
			expectedResult: [][]int{{-1,-1,2}, {-1,0,1}},
		},
		{
			name: "Test Case 2",
			nums: []int{0,1,1},
			expectedResult: [][]int{},
		},
		{
			name: "Test Case 3",
			nums: []int{0,0,0},
			expectedResult: [][]int{{0,0,0}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ThreeSum(test.nums) 
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("3Sum(%v) = %v, want = %v", test.nums, got, test.expectedResult)
			}
		})
	}
}