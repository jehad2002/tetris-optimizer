package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile(path string) ([]Tetromino, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tetros := []Tetromino{}
	temp := []string{}
	id := 'A'

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(temp) != 0 {
				t, err := parseTetromino(temp, id)
				if err != nil {
					return nil, fmt.Errorf("ERROR")
				}
				tetros = append(tetros, t)
				temp = []string{}
				id++
			}
			continue
		}
		temp = append(temp, line)
	}
	if len(temp) != 0 {
		t, err := parseTetromino(temp, id)
		if err != nil {
			return nil, fmt.Errorf("ERROR")
		}
		tetros = append(tetros, t)
	}
	return tetros, nil
}

func parseTetromino(lines []string, id rune) (Tetromino, error) {
	if len(lines) != 4 {
		return Tetromino{}, fmt.Errorf("ERROR")
	}

	positions := [][2]int{}
	for y, line := range lines {
		if len(line) != 4 {
			return Tetromino{}, fmt.Errorf("ERROR")
		}
		for x, ch := range line {
			if ch == '#' {
				positions = append(positions, [2]int{y, x})
			} else if ch != '.' {
				return Tetromino{}, fmt.Errorf("ERROR")
			}
		}
	}
	if len(positions) != 4 {
		return Tetromino{}, fmt.Errorf("ERROR")
	}

	minY, minX := positions[0][0], positions[0][1]
	for _, p := range positions {
		if p[0] < minY {
			minY = p[0]
		}
		if p[1] < minX {
			minX = p[1]
		}
	}
	t := Tetromino{id: id}
	for i := 0; i < 4; i++ {
		t.blocks[i][0] = positions[i][0] - minY
		t.blocks[i][1] = positions[i][1] - minX
	}
	return t, nil
}
