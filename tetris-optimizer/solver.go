package main

import "math"

func Solve(tetros []Tetromino) [][]rune {
    size := int(math.Ceil(math.Sqrt(float64(len(tetros) * 4))))
    for {
        grid := createGrid(size)
        if backtrack(grid, tetros, 0) {
            return grid
        }
        size++
    }
}

func backtrack(grid [][]rune, tetros []Tetromino, index int) bool {
    if index == len(tetros) {
        return true
    }

    t := tetros[index]
    size := len(grid)
    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            if canPlace(grid, t, y, x) {
                place(grid, t, y, x, t.id)
                if backtrack(grid, tetros, index+1) {
                    return true
                }
                place(grid, t, y, x, '.')
            }
        }
    }
    return false
}

func canPlace(grid [][]rune, t Tetromino, y, x int) bool {
    size := len(grid)
    for _, b := range t.blocks {
        ny, nx := y+b[0], x+b[1]
        if ny < 0 || ny >= size || nx < 0 || nx >= size || grid[ny][nx] != '.' {
            return false
        }
    }
    return true
}

func place(grid [][]rune, t Tetromino, y, x int, ch rune) {
    for _, b := range t.blocks {
        ny, nx := y+b[0], x+b[1]
        grid[ny][nx] = ch
    }
}