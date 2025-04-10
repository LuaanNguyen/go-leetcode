package valid_palindrome_II

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name 	string 
		s 	string 
		expectedResult 	bool 
	} {
		{
			name: "Test Case 1", 
			s: "aba",
			expectedResult: true, 
		},
		{
			name: "Test Case 2", 
			s: "abca",
			expectedResult: true, 
		},
		{
			name: "Test Case 3", 
			s: "abc",
			expectedResult: false, 
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := validPalindrome(test.s)

			if got != test.expectedResult {
				t.Errorf("checkPalindrome(%s) = %t; want = %t", test.s, got, test.expectedResult)
			}
		})
	}
}