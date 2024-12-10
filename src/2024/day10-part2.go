package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  file, err := os.Open("src/2024/day10.test")
  if err != nil {
    log.Fatalf("Error opening file %s", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var grid [][]int

  for scanner.Scan() {
    line := strings.Split(strings.TrimSpace(scanner.Text()), "")
    var row []int
    for _, r := range line {
      x, err := strconv.Atoi(r)
      if err != nil {
        log.Fatalf("Error casting string %s: %s", r, err)
        continue
      }
      row = append(row, x)
    }
    grid = append(grid, row)
  }

  ans := 0
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[0]); j++ {
      if grid[i][j] == 0 {
        x := getNumberOfTailheadsRatings(grid, i, j, -1)
        ans += x
      }
    }
  }
  fmt.Println(ans)
}

func getNumberOfTailheadsRatings(grid [][]int, i, j, prev int) int {
  if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
    return 0
  }

  if grid[i][j] - prev != 1 {
    return 0
  }

  if grid[i][j] == 9 {
    return 1
  }

  return getNumberOfTailheadsRatings(grid, i - 1, j, grid[i][j]) + getNumberOfTailheadsRatings(grid, i + 1, j, grid[i][j]) + getNumberOfTailheadsRatings(grid, i, j + 1, grid[i][j]) + getNumberOfTailheadsRatings(grid, i, j - 1, grid[i][j])
}
