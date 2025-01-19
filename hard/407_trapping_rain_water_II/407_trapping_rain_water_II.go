package trapping_rain_water_II

import "container/heap"

type Cell struct {
    Height  int
    Row     int
    Col     int
}

type MinHeap []Cell 

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
    if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

    ROWS := len(heightMap)
    COLS := len(heightMap[0])
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c ++ {
            if (r == 0 || r == ROWS - 1) || (c == 0 || c == COLS - 1) {
                heap.Push(minHeap, Cell{
                    Height: heightMap[r][c],
                    Row: r,
                    Col: c,
                })
                heightMap[r][c] = -1
            }
        }
    }

    res := 0
    maxHeight := 0 

    directions := [][]int {
        {0, 1},
        {0, -1},
        {1, 0},
        {-1, 0},
    }

    for minHeap.Len() > 0 {
        cell := heap.Pop(minHeap).(Cell)
        maxHeight = max(maxHeight, cell.Height)
        for _, dir := range directions {
            nr, nc := cell.Row + dir[0], cell.Col + dir[1]
            if nr >= 0 && nc >= 0 && nr < ROWS && nc < COLS && heightMap[nr][nc] != -1 {
                res += max(0, maxHeight - heightMap[nr][nc])
                heap.Push(minHeap,Cell{
                    Height: heightMap[nr][nc],
                    Row: nr,
                    Col: nc,
                }) 
                heightMap[nr][nc] = -1
            }
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}