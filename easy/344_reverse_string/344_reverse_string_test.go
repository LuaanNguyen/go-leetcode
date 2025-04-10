package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name 	string 
		s 	[]byte
		expectedResult 	[]byte 
	} {
		{
			name: "Test Case 1", 
			s: []byte{'h','e','l','l','o'},
			expectedResult: []byte{'o','l','l','e','h'}, 
		},
		{
			name: "Test Case 1", 
			s: []byte{'H','a','n','n','a','h'},
			expectedResult: []byte{'h','a','n','n','a','H'}, 
		},

		
	
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reverseString(test.s)
			if string(test.s) != string(test.expectedResult) {
				t.Errorf("reverseString(%s) = %s; want = %s", test.s, test.s, test.expectedResult)
			}
		})
	}
}