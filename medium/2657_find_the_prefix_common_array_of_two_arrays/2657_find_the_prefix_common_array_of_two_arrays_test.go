package find_the_prefix_common_array_of_two_arrays

import (
	"reflect"
	"testing"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	tests := []struct {
		name	string
		A	[]int
		B	[]int 
		expectedResult []int 
	} {
		{
			name: "Test Case 1",
			A : []int{1,3,2,4},
			B : []int{3,1,2,4},
			expectedResult: []int{0,2,3,4},
		},
		{
			name: "Test Case 2",
			A : []int{2,3,1},
			B : []int{3,1,2},
			expectedResult: []int{0,1,3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findThePrefixCommonArray(test.A, test.B)

			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("findThePrefixCommonArray(%v, %v) = %v; want = %v", test.A, test.B, got, test.expectedResult)
			}
		})
	}
}
	
	