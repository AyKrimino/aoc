package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("src/2024/day6.in")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	x, y := getStartingCoordinates(grid)
	grid[x][y] = '.'

	ans := 0
	for i, char := range grid {
		for j := range char {
			if i != x && j != y && grid[i][j] == '.' {
				grid[i][j] = '#'
				if hasCycle(grid, x, y) {
					ans++
				}
				grid[i][j] = '.'
			}
		}
	}

	fmt.Println(ans)
}

func getStartingCoordinates(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, char := range row {
			if char == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func hasCycle(grid [][]rune, x, y int) bool {
	directions := [4][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Bottom
		{0, -1}, // Left
	}
	currDirectionIdx := 0 // Default direction is Up
	limit := len(grid)*len(grid[0])*4 + 1
	i := 0
	for true {
		i++
		if i == limit {
			return true
		}
		nextX := x + directions[currDirectionIdx][0]
		nextY := y + directions[currDirectionIdx][1]
		if !isInsideGrid(nextX, nextY, len(grid), len(grid[0])) {
			return false
		}
		if grid[nextX][nextY] == '.' {
			x, y = nextX, nextY
		} else {
			currDirectionIdx = (currDirectionIdx + 1) % 4
		}
	}
	return false
}

func isInsideGrid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}
