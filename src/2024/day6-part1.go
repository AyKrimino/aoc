package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
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

  ans := getNumberOfVisitedPositions(grid, x, y)

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

func getNumberOfVisitedPositions(grid [][]rune, x, y int) int {
  visitedPositions := 0
  directions := [4][2]int{
    {-1, 0}, // Up
    {0, 1}, // Right
    {1, 0}, // Bottom
    {0, -1}, // Left
  }
  currDirectionIdx := 0 // Default direction is Up
  leftMap := false
  for !leftMap {
    if grid[x][y] != 'X' {
      visitedPositions++
      grid[x][y] = 'X'
    }
    nextX := x + directions[currDirectionIdx][0]
    nextY := y + directions[currDirectionIdx][1]
    leftMap = !isInsideGrid(nextX, nextY, len(grid), len(grid[0]))
    if !leftMap {
      if grid[nextX][nextY] == '#' {
        if currDirectionIdx + 1 == len(directions) {
          currDirectionIdx = -1
        }
        currDirectionIdx += 1
      }
      x += directions[currDirectionIdx][0]
      y += directions[currDirectionIdx][1]
    }
  }
  return visitedPositions
}

func isInsideGrid(x, y, rows, cols int) bool {
  return x >= 0 && x < rows && y >= 0 && y < cols
}
