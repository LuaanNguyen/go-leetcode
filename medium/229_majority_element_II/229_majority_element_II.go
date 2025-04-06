package majority_element_II

func majorityElement(nums []int) []int {
    count := make(map[int]int)

    for _, n := range nums {
        count[n]++ 

        if len(count) <= 2 {
            continue 
        }

        new_count := make(map[int]int)
        for k, v := range count {
            if v > 1 {
                new_count[k] = v - 1
            }
        }
        count = new_count
    }

    res := make([]int, 0)
    seen := make(map[int]bool)

	for k := range count {
		if seen[k] {
			continue
		}
		occurrence := 0
		for _, n := range nums {
			if n == k {
				occurrence++
			}
		}
		if occurrence > len(nums)/3 {
	        res = append(res, k)
		}
		seen[k] = true
	}

	return res
}