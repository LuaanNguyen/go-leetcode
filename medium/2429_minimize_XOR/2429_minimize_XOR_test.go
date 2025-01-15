package minimize_XOR

import "testing"

func TestMinimizeXOR(t *testing.T) {
	tests := []struct {
		name	string
		num1	int
		num2	int
		expectedResult 	int
	} {
		{
			name: "Test Case 1",
			num1: 3,
			num2: 5,
			expectedResult: 3,
		},
		{
			name: "Test Case 2",
			num1: 1,
			num2: 12,
			expectedResult: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t * testing.T) {
			got := minimizeXor(test.num1, test.num2)
			if got != test.expectedResult {
				t.Errorf("minimizeXOR(%d, %d) = %d; want = %d", test.num1, test.num2, got, test.expectedResult)
			}
	})
	}
}