package count_prefix_and_sufix_pairs_I

import "strings"

func countPrefixSuffixPairs(words []string) int {
    res := 0 

    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                res++
            }
        }
    }

    return res
}

func isPrefixAndSuffix(str1 string, str2 string) bool {
    return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}