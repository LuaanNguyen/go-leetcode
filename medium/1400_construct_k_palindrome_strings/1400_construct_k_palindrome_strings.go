package construct_k_palindrome_strings

func canConstruct(s string, k int) bool {
	//base case 
	if len(s) == k {
		return true 
	}

	if len(s) < k {
		return false
	}

	//frequency counter 
	frequency_counter := make(map[rune]int)
	for _, c := range s {
		frequency_counter[c]++ 
	}

	//count odd instances 
	odd_instances := 0 
	for _, val := range frequency_counter {
		if val % 2 == 1 {
			odd_instances ++ 
		}
	}

	//if the number of odd-frequency characters is greater than k, it's impossible to form a palindrome 
	return odd_instances <= k
}