package main

func createGrid(size int) [][]rune {
    grid := make([][]rune, size)
    for i := range grid {
        grid[i] = make([]rune, size)
        for j := range grid[i] {
            grid[i][j] = '.'
        }
    }
    return grid
}

func printGrid(grid [][]rune) {
    for _, row := range grid {
        for _, cell := range row {
            print(string(cell))
        }
        println()
    }
}