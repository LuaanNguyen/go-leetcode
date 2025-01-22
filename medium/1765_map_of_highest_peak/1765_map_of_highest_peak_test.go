package map_of_highest_peak

import (
	"testing"
)

func TestHighestPeek(t *testing.T) {
	tests := []struct {
		name	string 
		isWater 	[][]int
		expectedResult 	[][]int
	} {
		{
			name: "Test Case 1",
			isWater: [][]int {
				{0,1},
				{0,0},
			},
			expectedResult: [][]int {
				{1,0},
				{2,1},
			},
		},
		{
			name: "Test Case 2",
			isWater: [][]int {
				{0,0,1},
				{1,0,0},
				{0,0,0},
			},
			expectedResult: [][]int {
				{1,1,0},
				{0,1,1},
				{1,2,2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := highestPeak(test.isWater)
			if !equal2DSlices(got, test.expectedResult) {
				t.Errorf("highestPeek(%v) = %v; want = %v", test.isWater, got ,test.expectedResult)
			}
		})
	}
}

func equal2DSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}