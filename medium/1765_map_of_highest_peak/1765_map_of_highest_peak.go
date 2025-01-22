package map_of_highest_peak

func highestPeak(isWater [][]int) [][]int {
    rows := len(isWater)
    cols := len(isWater[0])
    directions := [][2]int{
        {1, 0}, {-1, 0}, {0, 1}, {0, -1},
    }

    // Initialize result grid with -1 for land and 0 for water
    res := make([][]int, rows)
    for i := range res {
        res[i] = make([]int, cols)
        for j := range res[i] {
            res[i][j] = -1
        }
    }

    // Initialize queue for BFS
    type point struct{ r, c int }
    q := []point{}

    // Add all water cells to the queue and mark them as height 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if isWater[r][c] == 1 {
                res[r][c] = 0
                q = append(q, point{r, c})
            }
        }
    }

    // Perform BFS
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        for _, dir := range directions {
            nr, nc := curr.r+dir[0], curr.c+dir[1]
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && res[nr][nc] == -1 {
                res[nr][nc] = res[curr.r][curr.c] + 1
                q = append(q, point{nr, nc})
            }
        }
    }

    return res
}
