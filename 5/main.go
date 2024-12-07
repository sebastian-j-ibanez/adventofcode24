package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const input = "input.txt"

// 1. Construct page order rule list
func getPageOrder() []int {
	var pageOrder []int

	return pageOrder
}

// 2. Get middle pages
func getUpdates() [][]int {
	var updates [][]int

	contents, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	unformattedUpdates := strings.Split(string(contents), "\n\n")
	for i, unformattedUpdate := range unformattedUpdates {
		updates[i] = append(updates[i], unformattedUpdate)
	}

	return updates
}

// 3. Sum middle pages
func middlePageSum(pageOrder []int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		for _, page := range update {
			index := slices.IndexFunc(pageOrder, func(n int) bool {
				return n == page
			})

			if index != -1 {
				sum += pageOrder[index]
			}
		}
	}

	return sum
}

func main() {
	pageOrder := getPageOrder()
	updates := getUpdates()
	fmt.Printf("Sum of middle page numbers: %d", middlePageSum(pageOrder, updates))
}
