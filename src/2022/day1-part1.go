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

	maxCalories := math.MinInt
	currCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maxCalories = max(maxCalories, currCalories)
			currCalories = 0
			continue
		}

		foodCalory, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error converting line to integer: %v", err)
		}
		currCalories += foodCalory
	}

	fmt.Println(maxCalories)
}
