package grid_game

import "testing"


func TestGridGame(t *testing.T) {
	tests := []struct {
		name 	string
		grid	[][]int
		expectedResult	int
	} {
		{
			name: "Test Case 1",
			grid: [][]int {
				{2,5,4},
				{1,5,1},
			},
			expectedResult: 4,
		},
		{
			name: "Test Case 2",
			grid: [][]int {
				{3,3,1},
				{8,5,2},
			},
			expectedResult: 4,
		},
		{
			name: "Test Case 3",
			grid: [][]int {
				{1,3,1,15},
				{1,3,3,1},
			},
			expectedResult: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			got := gridGame(test.grid)
			if got != test.expectedResult {
				t.Errorf("gridGame(%v) = %d; want = %d", test.grid, got, test.expectedResult)
			}
		})
	}

}