package is_subsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s	string
		t	string
		expectedResult 	bool 
	} {
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,	
		},
	}

	//Loop through test cases
	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			got := isSubsequence(test.s, test.t)

			if got != test.expectedResult {
				t.Errorf("For input s = %s and t = %s, expected result %t, but got %t", test.s, test.t, test.expectedResult, got)
			} 
		})
	}
}