package first_completely_painted_row_or_column

func firstCompleteIndex(arr []int, mat [][]int) int {
    ROWS := len(mat)
    COLS := len(mat[0])

    rowCount := make([]int, ROWS)
    colCount := make([]int, COLS)

    numToPos := make(map[int][2]int)

    for r := range ROWS {
        for c := range COLS {
            numToPos[mat[r][c]] = [2]int{r, c}
        }
    }

    for i := range len(arr) {
        pos := numToPos[arr[i]]
        r, c := pos[0], pos[1]

        rowCount[r]++
        colCount[c]++

        if rowCount[r] == COLS || colCount[c] == ROWS {
            return i
        }
    }

    return -1
}