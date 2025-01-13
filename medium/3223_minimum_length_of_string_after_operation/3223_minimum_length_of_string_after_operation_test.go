package minimum_length_of_string_after_operation

import "testing"


func TestMinimumLength(t *testing.T) {
	tests := []struct {
		name	string
		s	string
		expectedResult int 
	} {
		{
			name: "Test Case 1",
			s: "abaacbcbb",
			expectedResult: 5,
		},
		{
			name: "Test Case 2",
			s: "aa",
			expectedResult: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			got := minimumLength(test.s)
			if got != test.expectedResult {
				t.Errorf("minimumLength(\"%s\") = %d; want %d", test.s, got, test.expectedResult)
			}
	})
	}
}