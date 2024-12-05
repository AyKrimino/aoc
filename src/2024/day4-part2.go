package main

import (
  "fmt"
  "os"
  "log"
  "strings"
)

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
      if isValidXMAS(grid, i, j) {
        result += 1
      }
    }
  }
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

func isValidXMAS(grid [][]rune, i int, j int) bool {
  if !inBoundries(i, j, len(grid), len(grid[0])) {
    return false
  }

  if grid[i][j] != 'A' {
    return false
  }
  
  diag1Valid := (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || 
  (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')


  diag2Valid := (grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') || 
  (grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M')

  return diag1Valid && diag2Valid
}

func inBoundries(i int, j int, rows int, cols int) bool {
  return i - 1 >= 0 && i + 1 < rows && j - 1 >= 0 && j + 1 < cols
}
