package merge_intervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	res := [][]int{}

	// Sort by starting index in each element 
	sort.Slice(intervals, func(i, j int ) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0 
	for i < n {
		start := intervals[i][0] 
		end := intervals[i][1]
		j := i + 1

		for j < n && intervals[j][0] <= end {
			if end < intervals[j][1] {
				end = intervals[j][1]
			}
			j ++ 
		}   

		res = append(res, []int{start, end})
		i = j
	}

	return res
}