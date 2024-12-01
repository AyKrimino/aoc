package main

import (
  "fmt"
  "bufio"
  "os"
  "sort"
  "math"
)

func main() {
  file, err := os.Open("day1.in")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  var leftList []int
  var rightList []int

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    var left, right int
    _, err := fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
    if err != nil {
      fmt.Println("Error reading nubmers:", err)
      continue
    }
    leftList = append(leftList, left)
    rightList = append(rightList, right)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Scanner error:", err)
  }

  // sort leftList and rightList in ascending order
  sort.Slice(leftList, func(i, j int) bool {
    return leftList[i] < leftList[j]
  })
  sort.Slice(rightList, func(i, j int) bool {
    return rightList[i] < rightList[j]
  })

  diff := 0
  for i := range leftList {
    diff += int(math.Abs(float64(leftList[i] - rightList[i])))
  }
  
  fmt.Println(diff)
}
