package minimum_operations_to_make_array_values_equal_to_k

func minOperations(nums []int, k int) int {
    distinct := make(map[int]bool)

    for _, n := range nums {
        if n < k {
            return -1
        } else if n > k {
            distinct[n] = true
        }
    }
    return len(distinct)
}
