package main

import (
  "math"
  "fmt"
  "bufio"
  "strings"
  "os"
  "strconv"
)

func makeMeIntSplice(lst []string) ([]int, error) {
  res := make([]int, len(lst))
  for i, str := range lst {
    val, err := strconv.Atoi(str)
    if err != nil {
      return nil, fmt.Errorf("Failed to convert %q: %v", str, err)
    }
    res[i] = val
  }
  return res, nil
}

func isSafeReport(levels []int) bool {
    prevDiff := 0
    for i := 0; i < len(levels)-1; i++ {
        currDiff := levels[i] - levels[i+1]

        // Check for increasing/decreasing constraints
        if currDiff*prevDiff < 0 {
            return false
        }

        // Check adjacent level constraints
        diffAbs := math.Abs(float64(currDiff))
        if diffAbs < 1 || diffAbs > 3 {
            return false
        }
        prevDiff = currDiff
    }
    return true
}


func main() {
  file, err := os.Open("day2.in")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  ans := 0
  for scanner.Scan() {
    report := scanner.Text()
    levels, err := makeMeIntSplice(strings.Split(report, " "))
    if err != nil {
      fmt.Println("Skipping line due to error:", err)
      continue
    }
    
    if isSafeReport(levels){
      ans++
    } else {
      for i := 0; i < len(levels); i++ {
        levelsUpdated := append([]int{}, levels[:i]...)
        levelsUpdated = append(levelsUpdated, levels[i + 1:]...)
        if (isSafeReport(levelsUpdated)) {
          ans++
          break
        }
      }
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("scanner error:", err)
  }
  fmt.Println(ans)
}
