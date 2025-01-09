package counting_words_with_a_given_string

import "testing"

func TestPrefixCount(t *testing.T) {
	tests := []struct{
		words []string
		pref  string 
		expectedResult int
	}{
		{
			words:          []string{"pay", "attention", "practice", "attend"},
			pref:           "at",
			expectedResult: 2,
		},
		{
			words:          []string{"leetcode","win","loops","success"},
			pref:           "code",
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
            got := prefixCount(test.words, test.pref)
            if got != test.expectedResult {
				t.Errorf("expected %d, but got %d", test.expectedResult, got)
            }
        })

	}

}	