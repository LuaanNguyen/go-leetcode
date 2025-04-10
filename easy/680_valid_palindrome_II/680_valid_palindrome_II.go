package valid_palindrome_II

func validPalindrome(s string) bool {
    l := 0 
    r := len(s) - 1
    
    for l <= r {
        if s[l] != s[r] {
           return CheckPalindrome(s, l, r - 1) || CheckPalindrome(s,l + 1, r)
        }
        l ++ 
        r -- 
    }

    return true 
}

// Helper
func CheckPalindrome(s string, i, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i ++
        j --
    }
    return true
}