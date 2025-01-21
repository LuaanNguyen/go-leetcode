package first_completely_painted_row_or_column

import "testing"


func TestFirstCompleteIndex(t *testing.T) {
	tests := []struct {
		name 	string 
		arr		[]int
		mat		[][]int
		expectedResult	int
	} {
		{
			name: "Test Case 1",
			arr: []int{1,3,4,2},
			mat: [][]int{
				{1,4},
				{2,3},
			},
			expectedResult: 2,
		},
		{
			name: "Test Case 2",
			arr: []int{2,8,7,4,1,3,5,6,9},
			mat: [][]int{
				{3,2,5},
				{1,4,6},
				{8,7,9},
			},
			expectedResult: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := firstCompleteIndex(test.arr, test.mat)
			if got != test.expectedResult {
				t.Errorf("firstCompleteIndex(%v, %v) = %d; want = %d", test.arr, test.mat, got, test.expectedResult)
			}
		})
	}
}