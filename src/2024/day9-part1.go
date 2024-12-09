package main

import (
  "fmt"
  "os"
  "log"
  "strings"
)

func main() {
  bytes, err := os.ReadFile("src/2024/day9.in")
  if err != nil {
    log.Fatalf("Error reading the file %s\n", err)
  }
  input := strings.TrimSpace(string(bytes))

  var diskMap []int 

  x := 0
  for i := range input {
    v := int(input[i] - '0')
    if err != nil {
      log.Fatalf("Error converting rune %c to int: %s", v, err)
    }
    if i % 2 == 0 { // File
      for j := 0; j < v; j++ {
        diskMap = append(diskMap, x) 
      }
      x++
    } else { // Free space
      for j := 0; j < v; j++ {
        diskMap = append(diskMap, -1)
      }
    }
  }

  left, right := 0, len(diskMap) - 1
  for left < right {
    for diskMap[left] != -1 {
      left++
    }
    for diskMap[right] == -1 {
      right--
    }
    tmp := diskMap[left]
    diskMap[left] = diskMap[right]
    diskMap[right] = tmp
    right--
    left++
  }

  ans := 0
  for i, v := range diskMap {
    if v != -1 {
      ans += i * v
    }
  }

  fmt.Println("sum =", ans)
}
