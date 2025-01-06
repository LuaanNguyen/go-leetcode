package two_sum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, found := m[target-num]; found {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
