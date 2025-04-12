package four_sum

import (
	"reflect"
	"testing"
)

type Problem struct {
	name	string 
	nums	[]int
	target	int
	expectedResult 	[][]int
}

func TestFourSum(t *testing.T) {
	tests := []Problem  {
		{
			name: "Test Case 1",
			nums: []int{1,0,-1,0,-2,2},
			target: 0,
			expectedResult: [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}},
		},
		{
			name: "Test Case 1",
			nums: []int{1,0,-1,0,-2,2},
			target: 0,
			expectedResult: [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}},
		},
	}

	for  _, test := range tests {
		got := fourSum(test.nums, test.target) 

		if !reflect.DeepEqual(got, test.expectedResult) {
			t.Errorf("fourSum(%v, %d) = %v, want = %v", test.nums, test.target, got, test.expectedResult)
		}
	}
}