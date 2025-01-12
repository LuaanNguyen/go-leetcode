package check_if_a_parentheses_string_can_be_valid

import "testing"


func TestCanBeValid(t *testing.T) {
	tests := []struct {
		name	string
		s 	string 
		locked	string 
		expectedResult	bool
	} {
		{
			name: "Test Case 1",
			s: "))()))",
			locked: "010100",
			expectedResult: true,
		},
		{
			name: "Test Case 2",
			s: "()()",
			locked: "010100",
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t * testing.T) {
			got := canBeValid(test.s, test.locked)
			if got != test.expectedResult {
				t.Errorf("canBeValidConstruct(\"%s\", %s) = %t; want %t", test.s, test.locked, got, test.expectedResult)
			}
	})
	}
}