package find_the_prefix_common_array_of_two_arrays

func findThePrefixCommonArray(A []int, B []int) []int {
	A_count := make(map[int]bool)
	B_count := make(map[int]bool)

	res := make([]int, len(A))

	for i := 0; i < len(A); i ++ {
		A_count[A[i]] = true
		B_count[B[i]] = true

		for key := range A_count {
			if B_count[key] {
				res[i]++
			}
		}
	}
	return res 
}