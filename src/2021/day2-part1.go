package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  file, err := os.Open("day2.in")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  horizontalPosition := 0
  depth := 0

  for scanner.Scan() {
    var command string
    var units int
    _, err := fmt.Sscanf(scanner.Text(), "%s %d", &command, &units)
    if err != nil {
      fmt.Println("Error reading from file:", err)
      continue
    }
    
    if command == "forward" {
      horizontalPosition += units
    } else if command == "down" {
      depth += units
    } else {
      depth -= units
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Scanner error:", err)
  }
  fmt.Println(horizontalPosition * depth)
}
