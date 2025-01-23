package count_servers_that_communicate

import "testing"


func TestCountServer(t *testing.T) {
	tests := []struct {
		name 	string 
		grid 	[][]int
		expectedResult	int
	} {
		{
			name: "Test Case 1",
			grid: [][]int {
				{1,0},
				{0,1},
			},
			expectedResult: 0,
		},
		{
			name: "Test Case 2",
			grid: [][]int {
				{1,0},
				{1,1},
			},
			expectedResult: 3,
		},
		{
			name: "Test Case 3",
			grid: [][]int {
				{1,1,0,0},
				{0,0,1,0},
				{0,0,1,0},
				{0,0,0,1},
			},
			expectedResult: 4,
		},
	}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := countServers(test.grid)
			if got != test.expectedResult {
				t.Errorf("countServers(%v) = %d; want = %d", test.grid, got, test.expectedResult)
			}
		})
	}
}