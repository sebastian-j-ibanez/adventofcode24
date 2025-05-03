package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

const input = "input.txt"

// Get the page updates and the page order rules
// from the input file.
func getFileData() ([][]int, map[int][]int) {
	contents, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// fileData[0] -> page order rules
	// fileData[1] -> updates
	fileData := strings.Split(string(contents), "\n\n")
	pageOrderLines := strings.Split(fileData[0], "\n")

	// Get page order rules (each rule is an int slice ex. [86, 92])
	var pageOrder [][]int
	for _, rule := range pageOrderLines {
		strPages := strings.Split(rule, "|")
		pageOrder = append(pageOrder, pageNumbersToInt(strPages))
	}

	rules := rulesToMap(pageOrder)

	// Get page updates
	var updates [][]int
	updateLines := strings.Split(fileData[1], "\n")
	updateLines = updateLines[:len(updateLines)-1]
	for _, line := range updateLines {
		update := strings.Split(line, ",")
		updates = append(updates, updatesToInt(update))
	}
	return updates, rules
}

// Get sum of all middle pages in valid updates.
func middlePageSum(updates [][]int, pageOrderRules map[int][]int) int {
	sum := 0

	for _, update := range updates {
		if validUpdate(update, pageOrderRules) {
			index := int(math.Floor(float64(len(update)) / float64(2))) // Middle index of update (it looks nasty, I know)
			sum += update[index]
		}
	}

	return sum
}

// Return true if update follows page order rules.
func validUpdate(update []int, pageOrderRules map[int][]int) bool {
	var prevPages []int
	for _, page := range update {
		if pageOrderRules[page] != nil {
			for _, prevPage := range prevPages {
				if slices.Contains(pageOrderRules[page], prevPage) {
					return false
				}
			}
		}
		prevPages = append(prevPages, page)
	}

	return true
}

// PART 2
// Reorder numbers in updates to meet the page order rules.
func fixUpdates(updates [][]int, pageOrderRules map[int][]int) [][]int {
	var fixedUpdates [][]int

	for _, update := range updates {
		if !validUpdate(update, pageOrderRules) {
			var prevPages []int
			for i, page := range update {
				for _, prevPage := range prevPages { 
					if pageOrderRules[page] != nil &&
						slices.Contains(pageOrderRules[page], prevPage) {
						for j, thing := range update {
							if thing == prevPage {
								tmp := thing
								update[j] = update[i]
								update[i] = tmp
							}
						}
					}
				}
				prevPages = append(prevPages, page)
			}
		}
		fixedUpdates = append(fixedUpdates, update)
	}

	return fixedUpdates
}

func main() {
	updates, pageOrder := getFileData()
	part1Sum := middlePageSum(updates, pageOrder)
	updates = fixUpdates(updates, pageOrder)
	part2Sum := middlePageSum(updates, pageOrder)

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
}
