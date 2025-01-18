package neighboring_bitwise_XOR

import "testing"


func TestDoesValidArrayExist(t *testing.T) {
	tests := []struct {
		name	string 
		derived	[]int
		expectedResult	bool
	} {
		{
			name: "Test Case 1",
			derived: []int{1,1,0},
			expectedResult: true,
		},
		{
			name: "Test Case 2",
			derived: []int{1,1},
			expectedResult: true,
		},
		{
			name: "Test Case 3",
			derived: []int{1,0},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := doesValidArrayExist(test.derived)
			if got != test.expectedResult {
				t.Errorf("doesValidArrayExist(%v) = %t; want = %t", test.derived, got, test.expectedResult)
			}
		})
	}
}