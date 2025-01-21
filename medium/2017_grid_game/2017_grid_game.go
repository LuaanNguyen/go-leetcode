package grid_game

import "math"

func gridGame(grid [][]int) int {
    firstRowSum := 0
    for _, value := range grid[0] {
        firstRowSum += value
    }

    secondRowSum := 0
    minimumSum := math.MaxInt

    for turnIndex := 0; turnIndex < len(grid[0]); turnIndex++ {
        firstRowSum -= grid[0][turnIndex]
        minimumSum = min(minimumSum, max(firstRowSum, secondRowSum))
        secondRowSum += grid[1][turnIndex]
    }

    return minimumSum
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}