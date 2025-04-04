package sort_colors

func SortColors(nums []int) []int  {
    low := 0 
    mid := 0 
    high := len(nums) - 1

    for mid <= high {
        if nums[mid] == 0 {
            temp := nums[mid]
            nums[mid] = nums[low]
            nums[low] = temp 

            mid++
            low++
        } else if nums[mid] == 1 {
            mid++
        } else if nums[mid] == 2 {
            temp := nums[mid]
            nums[mid] = nums[high]
            nums[high] = temp

            high--
        }
    }

	return nums
}