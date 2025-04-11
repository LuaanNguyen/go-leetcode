package count_symmetric_integers

import "testing"

func TestCountSymmetricIntegers(t *testing.T) {
	tests := []struct {
		name	string 
		low	 int
		high 	int
		expectedResult 	int 
	} {
		{
			name: "Test Case 1", 
			low: 1,
			high: 100, 
			expectedResult: 9,
		},
		{
			name: "Test Case 2", 
			low: 1200,
			high: 1230, 
			expectedResult: 4,
		},	
	}

	for _, test := range tests {
		got := countSymmetricIntegers(test.low, test.high)
		if got != test.expectedResult {
			t.Errorf("countSymmetricIntegers(%d, %d) = %d; want = %d", test.low, test.high, got, test.expectedResult)
		}
	}
}