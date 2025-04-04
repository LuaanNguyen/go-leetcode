package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name	string
		nums	[]int 
		expectedResult []int
	} {
		{
			name: "Test Case 1",
			nums: []int{2,0,2,1,1,0},
			expectedResult: []int{0,0,1,1,2,2},
		}, 
		{
			name: "Test Case 2",
			nums: []int{2, 0 ,1},
			expectedResult: []int{0, 1, 2},
		}, 
	}

	for _, test := range tests {
		t.Run(test.name, func(t * testing.T) {
			got := SortColors(test.nums)
			if  !reflect.DeepEqual(got, test.expectedResult){
				t.Errorf("SortColors(%v) = %v; want = %v", test.nums, got, test.expectedResult)
			}
		})
	}
}