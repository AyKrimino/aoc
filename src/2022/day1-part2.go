package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.in")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var topThreeCalories [3]int
	currCalories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			minVal := math.MaxInt
			minPos := -1
			for i := 0; i < 3; i++ {
				if topThreeCalories[i] < minVal {
					minVal = topThreeCalories[i]
					minPos = i
				}
			}
			if minPos != -1 && topThreeCalories[minPos] < currCalories {
				topThreeCalories[minPos] = currCalories
			}
			currCalories = 0
			continue
		}

		foodCalory, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error converting line to integer: %v", err)
		}
		currCalories += foodCalory
	}

	totalCalories := 0
	for i := 0; i < 3; i++ {
		totalCalories += topThreeCalories[i]
	}

	fmt.Println(totalCalories)
}
