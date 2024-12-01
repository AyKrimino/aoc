package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  file, err := os.Open("day1.in")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  var leftList []int
  frequencyMap := make(map[int]int)
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    var left, right int
    _, err := fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
    if err != nil {
      fmt.Println("Error reading numbers:", err)
      continue
    }

    leftList = append(leftList, left)
    frequencyMap[right] += 1
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Scanner error:", err)
  }
  
  ans := 0
  for _, num := range leftList {
    ans += num * frequencyMap[num]
  }

  fmt.Println(ans)
}
