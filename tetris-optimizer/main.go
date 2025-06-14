package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("ERROR")
        return
    }
    tetros, err := ParseFile(os.Args[1])
    if err != nil {
        fmt.Println("ERROR")
        return
    }
    grid := Solve(tetros)
    printGrid(grid)
}