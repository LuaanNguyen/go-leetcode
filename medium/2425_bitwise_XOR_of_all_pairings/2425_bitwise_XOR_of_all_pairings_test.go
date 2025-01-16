package bitwise_XOR_of_all_pairings

import "testing"

func TestXorAllNums(t *testing.T) {
	tests := []struct {
		name	string 
		nums1	[]int
		nums2	[]int
		expectedResult	int
	} {
		{
			name: "Test Case 1",
			nums1: []int{2,1,3},
			nums2: []int{10,2,5,0},
			expectedResult: 13,
		},
		{
			name: "Test Case 2",
			nums1: []int{1,2},
			nums2: []int{3,4},
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := xorAllNums(test.nums1, test.nums2)

			if got != test.expectedResult {
				t.Errorf("xorAllNums(%v, %v) = %d; want = %d", test.nums1, test.nums2, got, test.expectedResult)
			}
		})
	}
}