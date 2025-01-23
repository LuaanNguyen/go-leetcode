package count_servers_that_communicate

func countServers(grid [][]int) int {
    ROWS := len(grid)
    COLS := len(grid[0])

    rowCount := make([]int, ROWS)
    colCount := make([]int, COLS)

    for r := range ROWS {
        for c := range COLS {
            if grid[r][c] == 1 {
                rowCount[r]++
                colCount[c]++
            }
        }
    }

    res := 0

    for r := range ROWS {
        for c := range COLS {
            if grid[r][c] == 1 && (rowCount[r] > 1 || colCount[c] > 1) {
                res++
            }
        }
    }

    return res
    
}