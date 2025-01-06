
package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
    // Define test cases
    tests := []struct {
        nums   []int
        target int
        result []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
    }

    // Loop through test cases
    for _, test := range tests {
        t.Run("Test", func(t *testing.T) {
            got := twoSum(test.nums, test.target)
            if len(got) != len(test.result) {
                t.Errorf("Expected result %v, but got %v", test.result, got)
            }

            // Compare the results
            for i := range got {
                if got[i] != test.result[i] {
                    t.Errorf("For input %v with target %d, expected result %v, but got %v", test.nums, test.target, test.result, got)
                }
            }
        })
    }
}
