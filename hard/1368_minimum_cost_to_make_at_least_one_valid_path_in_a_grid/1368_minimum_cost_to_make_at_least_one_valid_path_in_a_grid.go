package minimum_cost_to_make_at_least_one_valid_path_in_a_grid

import (
	"container/heap"
	"math"
)

func minCost(grid [][]int) int {
    directions := [][]int {
        {0, 1},
        {0, -1},
        {1, 0},
        {-1, 0},
    }

    ROWS := len(grid)
    COLS := len(grid[0])

    minCost := make([][]int, ROWS)

    for r := range minCost{
        minCost[r] = make([]int, COLS)
        for c := range minCost[r] {
            minCost[r][c] = math.MaxInt
        }
    }
    minCost[0][0] = 0 

    q := &PriorityQueue{}
    heap.Init(q)
    heap.Push(q, &Node{row: 0, col: 0, cost: 0})

    for q.Len() > 0 {
		current := heap.Pop(q).(*Node)
		row, col, currCost := current.row, current.col, current.cost

		if currCost > minCost[row][col] {
			continue
		}

		for dirIdx, dir := range directions {
			dr, dc := dir[0], dir[1]
			newRow, newCol := row+dr, col+dc

			cost := 0
			if grid[row][col] != dirIdx+1 {
				cost = 1
			}

			if newRow >= 0 && newRow < ROWS && newCol >= 0 && newCol < COLS &&
				currCost+cost < minCost[newRow][newCol] {
				minCost[newRow][newCol] = currCost + cost
				if cost == 1 {
					heap.Push(q, &Node{row: newRow, col: newCol, cost: currCost + cost})
				} else {
					heap.Push(q, &Node{row: newRow, col: newCol, cost: currCost})
				}
			}
		}
    }

    return minCost[ROWS -1][COLS -1]
}


type Node struct {
	row, col int
	cost     int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}