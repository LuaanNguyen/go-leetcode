package four_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums) // Sort arr in asc order 

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] { // remove duplicatefor i
            continue
        }
        for j := i + 1; j < len(nums); j++ {
            if j > i + 1 && nums[j] == nums[j - 1] { // remove duplicatefor j
                continue
            }
            l := j + 1
            r := len(nums) - 1
            for l < r {
                if nums[i] + nums[j] + nums[l] + nums[r] == target {
                    res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

                    for l < r && nums[l] == nums[l + 1] {
                        l ++ 
                    } 
                    for l < r && nums[r] == nums[r - 1] {
                        r --
                    }
                    l ++ 
                    r -- 
                } else if nums[i] + nums[j] + nums[l] + nums[r] > target {
                    r -- 
                } else {
                    l ++ 
                }
            }
        }
    }
    return res
}