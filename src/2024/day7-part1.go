package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, err := os.Open("src/2024/day7.in")
  if err != nil {
    log.Fatalf("Failed to open file: %s", err)
  }
  defer file.Close()

  ans := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Split(line, ": ")
    testValue, err := strconv.Atoi(parts[0])
    if err != nil {
      log.Fatalf("error %s", err)
    }
    var numbers []int
    for _, str := range strings.Split(parts[1], " ") {
      nb, err := strconv.Atoi(str)
      if err != nil {
        log.Fatalf("error %s", err)
      }
      numbers = append(numbers, nb)
    }

    if dfs(numbers, testValue, 0, numbers[0]) {
      ans += testValue
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatalf("Error reading file: %s", err)
  }

  fmt.Println(ans)
}

func dfs(numbers []int, target, i, currSum int) bool {
  if i == len(numbers)-1 {
    return currSum == target
  }
  if currSum > target {
    return false
  }
  return dfs(numbers, target, i+1, currSum+numbers[i+1]) || dfs(numbers, target, i+1, currSum*numbers[i+1])
}
