package two_sum_II_input_array_is_sorted

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name 	string
		numbers	[]int
		target	int
		expectedResult	[]int 
	} {
		{
			name: "Test Case 1",
			numbers: []int{2,7,11,15},
			target: 9,
			expectedResult: []int{1,2},
		},
		{
			name: "Test Case 2",
			numbers: []int{2,3,4},
			target: 6,
			expectedResult: []int{1,3},
		},
		{
			name: "Test Case 3",
			numbers: []int{-1,0},
			target: -1,
			expectedResult: []int{1,2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := twoSum(test.numbers, test.target)
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("twoSum(%v, %d) = %v ; want = %v", test.numbers, test.target, got, test.expectedResult)
			}
		})
	}
}