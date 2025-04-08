package three_sum

import "sort"

func ThreeSum(nums []int) [][]int {
    res := make([][]int, 0)

    //sort in asc 
    sort.Slice(nums, func(i , j int) bool {
        return nums[i] < nums[j]
    })

    for i := range nums {
        if i > 0 && nums[i] == nums[i - 1] {
            continue 
        }

        target := -nums[i]
        j := i + 1 
        k := len(nums) - 1

        for j < k {
            if target == nums[j] + nums[k] {
                res = append(res, []int{nums[i], nums[j], nums[k]})

                for j < k && nums[j] == nums[j + 1] {
                    j++ 
                }
                for j < k  && nums[k] == nums[k - 1] {
                    k--
                }

                j ++ 
                k -- 
            } else if target > nums[j] + nums[k] {
                j ++ 
            } else {
                k --
            }
        }
    }

    return res
}