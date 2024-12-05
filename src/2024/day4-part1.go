package main

import (
  "fmt"
  "os"
  "log"
  "strings"
)

var directions = [8][2]int{
	{0, -1},  
	{0, 1},  
	{-1, 0},
	{1, 0},
	{-1, -1},
  {-1, 1},
	{1, -1},
	{1, 1},
}

func main() {
  b, err := os.ReadFile("src/2024/day4.in")
  if err != nil {
    log.Fatal(err)
  }
  input := string(b)

  grid := convertToMatrix(input)
  result := 0
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
      result += getNumberOf(grid, i, j, "XMAS", [2]int{0, 0})
    }
  }
  fmt.Println(grid)
  fmt.Println(len(grid))
  fmt.Println(len(grid[0]))
  fmt.Println(result)
}

func convertToMatrix(input string) [][]rune {
  lines := strings.Split(input, "\n")
  matrix := make([][]rune, len(lines) - 1)
  
  for i, line := range lines {
    if line != "" {
      matrix[i] = []rune(line)
    }
  }
  return matrix
}




func getNumberOf(grid [][]rune, i int, j int, pattern string, direction [2]int) int {
  if !isInsideGrid(len(grid), len(grid[0]), i, j) {
    return 0
  }

  if grid[i][j] != rune(pattern[0]) {
    return 0
  }

  if len(pattern) == 1 {
    return 1
  }

  original := grid[i][j]
  grid[i][j] = '.'

  nbr := 0
  if direction == [2]int{0, 0} {
    for _, dir := range directions {
      x, y := i+dir[0], j+dir[1]
      nbr += getNumberOf(grid, x, y, pattern[1:], dir)
    }
  } else {
    x, y := i + direction[0], j + direction[1]
    nbr += getNumberOf(grid, x, y, pattern[1:], direction)
  }
  grid[i][j] = original
  return nbr
}



func isInsideGrid(rows, cols, i, j int) bool {
  return i >= 0 && i < cols && j >= 0 && j < cols
}
