package minimum_cost_to_make_at_least_one_valid_path_in_a_grid

import "testing"


func TestMinCost(t *testing.T) {
	tests := []struct {
		name	string
		grid 	[][]int
		expectedResult 	int
	} {
		{
			name: "Test Case 1",
			grid: [][]int {
				{1,1,1,1},
				{2,2,2,2},
				{1,1,1,1},
				{2,2,2,2},
			},
			expectedResult: 3,
		},
		{
			name: "Test Case 2",
			grid: [][]int {
				{1,1,3},
				{3,2,2},
				{1,1,4},
			},
			expectedResult: 0,
		},
		{
			name: "Test Case 3",
			grid: [][]int {
				{1,2},
				{4,3},
			},
			expectedResult: 1,
		},
	}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minCost(test.grid)
			if got != test.expectedResult {
				t.Errorf("minCost(%v) = %d; want = %d", test.grid, got, test.expectedResult)
			}
	})
	}
}