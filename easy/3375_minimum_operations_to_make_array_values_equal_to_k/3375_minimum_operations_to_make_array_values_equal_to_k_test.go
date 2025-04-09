package minimum_operations_to_make_array_values_equal_to_k

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name 	string 
		nums	[]int
		k 	int 
		expectedResult	int 
	} {
		{
			name: "Test Case 1", 
			nums: []int{5,2,5,4,5},
			k : 2, 
			expectedResult: 2,
		},
		{
			name: "Test Case 2", 
			nums: []int{2,1,2},
			k : 2, 
			expectedResult: -1,
		},
		{
			name: "Test Case 3", 
			nums: []int{9,7,5,3},
			k : 1, 
			expectedResult: 4,
		},
	}

	for _, test := range tests {
		got := minOperations(test.nums, test.k) 

		if got != test.expectedResult {
			t.Errorf("minOperations(%v, %d) = %d, want = %d", test.nums, test.k, got, test.expectedResult )
		}
	}
}