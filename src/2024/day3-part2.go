package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "strings"
)

func main() {
  b, err := os.ReadFile("src/2024/day3.in")
  if err != nil {
    log.Fatal(err)
  }

  memory := string(b)
  enabled := true // Multiplications are enabled at the start
  totalSum := 0

  i := 0
  for i < len(memory) {
    if strings.HasPrefix(memory[i:], "don't()") {
      enabled = false
      i += len("don't()")
      continue
    }
    if strings.HasPrefix(memory[i:], "do()") {
      enabled = true
      i += len("do()")
      continue
    }

    if enabled && strings.HasPrefix(memory[i:], "mul(") {
      closeIdx := strings.Index(memory[i:], ")")
      if closeIdx != -1 {
        content := memory[i+4 : i+closeIdx] 
        parts := strings.Split(content, ",")

        if isValidParts(parts) { 
          a, err1 := strconv.Atoi(parts[0])
          b, err2 := strconv.Atoi(parts[1])
          if err1 == nil && err2 == nil {
            totalSum += a * b
          }
        } else {
          i += 5
          continue
        }
        i += closeIdx + 1 
        continue
      }
    }
    i++
  }

  fmt.Println(totalSum)
}

func isValidParts(parts []string) bool {
	if len(parts) != 2 {
		return false
	}
	return isValidLength(parts[0]) && isValidLength(parts[1])
}

func isValidLength(part string) bool {
	return len(part) >= 1 && len(part) <= 3
}
