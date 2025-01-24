package find_eventual_safe_states

import "testing"


func TestEventualSafeNodes(t *testing.T) {
	tests := []struct {
		name 	string 
		graph 	[][]int
		expectedResult 	[]int 
	} {
		{
			name: "Test Case 1",
			graph: [][]int{
			{1,2},
			{2,3},
			{5},
			{0},
			{5},
			{},
			{},
		},
		expectedResult: []int{2,4,5,6},
		},
	}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := eventualSafeNodes(test.graph)
			if !equalSlice(got, test.expectedResult) {
				t.Errorf("eventual")
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}