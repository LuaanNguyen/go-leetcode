package majority_element_II

import (
	"reflect"
	"testing"
)

func MajorityElementTest(t *testing.T) {
	tests := []struct {
		name	string
		nums 	[]int
		expectedResult []int
	} {
		{
			name: "Test Case 1",
			nums: []int{3,2,3},
			expectedResult: []int{3},
		},
		{
			name: "Test Case 2",
			nums: []int{1},
			expectedResult: []int{1},
		},
		{
			name: "Test Case 3",
			nums: []int{1,2},
			expectedResult: []int{1,2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t * testing.T) {
			got := majorityElement(test.nums)
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("majorityElement( %v ) = %v; want = %v", test.nums, got, test.expectedResult)
			}
		})
	}
}