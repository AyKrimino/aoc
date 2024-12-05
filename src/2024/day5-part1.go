package main

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
  "slices"
)

func main() {
  b, err := os.ReadFile("src/2024/day5.in")
  if err != nil {
    log.Fatal(err)
  }
  input := string(b)
  lines := strings.Split(strings.TrimSpace(input), "\n")

  after := make(map[int][]int)
  before := make(map[int][]int)
  var updates [][]int
  secondPartReached := false

  for _, line := range lines {
    line = strings.TrimSpace(line)
    if line == "" {
      secondPartReached = true
      continue
    }

    if !secondPartReached {
      pages := strings.Split(line, "|")
      if len(pages) != 2 {
        fmt.Println("Skipping malformed line:", line)
        continue
      }

      x, err1 := strconv.Atoi(strings.TrimSpace(pages[0]))
      y, err2 := strconv.Atoi(strings.TrimSpace(pages[1]))
      if err1 != nil || err2 != nil {
        fmt.Printf("Error parsing line '%s': %v, %v\n", line, err1, err2)
        continue
      }

      after[y] = append(after[y], x)
      before[x] = append(before[x], y)

    } else {
      numStrs := strings.Split(line, ",")
      var tmp []int
      for _, val := range numStrs {
        num, err := strconv.Atoi(strings.TrimSpace(val))
        if err != nil {
          fmt.Printf("Error parsing number '%s' in line '%s': %v\n", val, line, err)
          continue
        }
        tmp = append(tmp, num)
      }
      updates = append(updates, tmp)
    }
  }

  ans := 0
  for _, update := range updates {
    if inCorrectOrder(update, before, after) {
      ans += getMid(update)
    } 
  }

  fmt.Println(ans)
}

func inCorrectOrder(list []int, before, after map[int][]int) bool {
  for i, li := range list {
    if !checkBeforeAll(list, li, i + 1, before) || !checkAfterAll(list, li, i - 1, after) {
      return false
    }
  }
  return true
}

func checkBeforeAll(list []int, page, startIdx int, before map[int][]int) bool {
  for i := startIdx; i < len(list); i++ {
    if !slices.Contains(before[page], list[i]) {
      return false
    }
  }
  return true
}

func checkAfterAll(list []int, page, endIdx int, after map[int][]int) bool {
  for i := 0; i <= endIdx; i++ {
    if !slices.Contains(after[page], list[i]) {
      return false
    }
  }
  return true
}

func getMid(list []int) int {
  return list[len(list) / 2]
}
