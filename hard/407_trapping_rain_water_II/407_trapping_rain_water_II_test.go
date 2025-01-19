package trapping_rain_water_II

import "testing"

func TestTrapRainWater(t *testing.T) {
	tests := []struct {
		name	string
		heightMap	[][]int
		expectedResult 	int
	} {
		{
			name: "Test Case 1",
			heightMap: [][]int {
				{1,4,3,1,3,2},
				{3,2,1,3,2,4},
				{2,3,3,2,3,1},
			},
			expectedResult: 4,
		},
		{
			name: "Test Case 2",
			heightMap: [][]int {
				{3,3,3,3,3},
				{3,2,2,2,3},
				{3,2,1,2,3},
				{3,2,2,2,3},
				{3,3,3,3,3},
			},
			expectedResult: 10,
		},
	}


	for _, test := range tests {
		t.Run(test.name, func (t * testing.T){
			got := trapRainWater(test.heightMap)
			if got != test.expectedResult {
				t.Errorf("trapRainWater(%v) = %d; want = %d\n", test.heightMap, got, test.expectedResult)
			}
		})
	}
}