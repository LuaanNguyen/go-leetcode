package bitwise_XOR_of_all_pairings

func xorAllNums(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)


	freq := make(map[int]int)

	for _, num := range nums1 {
		freq[num] += len2
	}

	for _, num := range nums2 {
		freq[num] += len1
	}

	res := 0 
	for num, count := range freq {
		if count % 2 != 0 {
			res ^= num
		}
	}

	return res
}