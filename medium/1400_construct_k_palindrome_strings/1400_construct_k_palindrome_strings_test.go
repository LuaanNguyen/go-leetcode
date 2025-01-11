package construct_k_palindrome_strings

import "testing"

func TestCanConstruct(t *testing.T){
	tests := []struct {
		name 	string 
		s		string 
		k 		int 
		expectedResult 	bool
	} {
		{
			name: "Test Case 1",
			s: "annabelle",
			k: 2,
			expectedResult: true,
		},
		{
			name: "Test Case 2",
			s: "leetcode",
			k: 3,
			expectedResult: false,
		},
		{
			name: "Test Case 3",
			s: "true",
			k: 4,
			expectedResult: true,
		},
		{
			name: "Test Case 3",
			s: "false",
			k: 4,
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t * testing.T) {
			got := canConstruct(test.s, test.k)
			if got != test.expectedResult {
				t.Errorf("canConstruct(\"%s\", %d) = %t; want %t", test.s, test.k, got, test.expectedResult)
			}
	})
	}
}