package count_prefix_and_sufix_pairs_I

import (
	"testing"
)

func TestCountPrefixSuffixPairs(t *testing.T) {
    tests := []struct {
        name     string
        words    []string
        expected int
    }{
        {
            name:     "Test case 1",
            words:    []string{"pa","papa","ma","mama"},
            expected: 2, 
        },
		{
            name:     "Test case 2",
            words:    []string{"abab","ab"},
            expected: 0, 
        },
       
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := countPrefixSuffixPairs(tt.words)
            if got != tt.expected {
                t.Errorf("countPrefixSuffixPairs(%v) = %v; want %v", tt.words, got, tt.expected)
            }
        })
    }
}
