package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const input = "input.txt"

// Part 1
func getTotalMulSum() int {
	// Read input file
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Parse regexp matches, sum matching products
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		if len(match) == 3 {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				sum += num1 * num2
			}
		}
	}

	return sum
}

// Part 2
func getEnabledMulSum() int {
	// Read input file
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Parse regexp matches, sum matching & enabled products
	sum := 0
	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	mulEnabled := true
	for _, match := range matches {
		switch {
		case match[1] == "do()":
			mulEnabled = true
		case match[1] == "don't()":
			mulEnabled = false
		case match[2] != "" && match[3] != "" && mulEnabled:
			num1, err1 := strconv.Atoi(match[2])
			num2, err2 := strconv.Atoi(match[3])
			if err1 == nil && err2 == nil {
				sum += num1 * num2
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("Sum of total multiplications: %d\n", getTotalMulSum())
	fmt.Printf("Sum of enabled multiplications: %d\n", getEnabledMulSum())
}
