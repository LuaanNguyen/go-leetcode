package two_sum_II_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
    l := 0
    r := len(numbers) - 1
    
    for {
        if numbers[l] + numbers[r] < target {
            l++
        } else if numbers[l] + numbers[r] > target {
            r--
        } else {
            break 
        }
    }
    return []int{l + 1, r + 1}
}