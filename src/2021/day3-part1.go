package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day3.in")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var ones []int
	var fileLength int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fileLength++

		for idx, val := range line {
			if idx >= len(ones) {
				ones = append(ones, 0)
			}
			if val == '1' {
				ones[idx]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	gamma_binary := getGammaValue(&ones, fileLength)
	epsilon_binary := getEpsilonValue(gamma_binary)

	gamma, err := strconv.ParseInt(gamma_binary, 2, 64)
	if err != nil {
		log.Fatalf("error converting %s to its integer value: %v", gamma_binary, err)
	}

	epsilon, err := strconv.ParseInt(epsilon_binary, 2, 64)
	if err != nil {
		log.Fatalf("error converting %s to its integer value: %v", epsilon_binary, err)
	}

	powerConsumption := gamma * epsilon

	fmt.Println(powerConsumption)
}

func getGammaValue(ones *[]int, fileLength int) string {
	binary_result := ""

	for _, one := range *ones {
		zero := fileLength - one
		if one > zero {
			binary_result += "1"
		} else {
			binary_result += "0"
		}
	}

	return binary_result
}

func getEpsilonValue(gamma_binary string) string {
	binary_result := ""

	for _, val := range gamma_binary {
		if val == '1' {
			binary_result += "0"
		} else {
			binary_result += "1"
		}
	}

	return binary_result
}
