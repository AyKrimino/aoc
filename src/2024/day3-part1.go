package main

import (
  "fmt"
  "os"
  "regexp"
  "strings"
  "strconv"
  "log"
)

func splitNumbers(instruction string) (int, int) {
  opening_parenthesis := strings.Index(instruction, "(")
  closing_parenthesis := strings.Index(instruction, ")")
  comma := strings.Index(instruction, ",")
  x, err := strconv.Atoi(instruction[opening_parenthesis + 1:comma])
  if err != nil {
    log.Fatal(err)
  }
  y, err := strconv.Atoi(instruction[comma + 1: closing_parenthesis])
  if err != nil {
    log.Fatal(err)
  }
  return x, y
}

func main() {
  bytesContent, err := os.ReadFile("day3.in")
  if err != nil {
    log.Fatal(err)
  }
  
  memory := string(bytesContent)

  regexPattern := `mul\(\d{1,3},\d{1,3}\)`
  re := regexp.MustCompile(regexPattern)

  instructions := re.FindAllString(memory, -1)
  ans := 0
  for _, instruction := range instructions {
    x, y := splitNumbers(instruction)
    ans += x * y
  }

  fmt.Println(ans)
}
