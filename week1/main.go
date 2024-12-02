package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	file = "input.txt"
)

func readFile() ([]int, []int) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	var listA, listB []int

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		numA, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		numB, err := strconv.Atoi(numbers[3])
		if err != nil {
			panic(err)
		}

		listA = append(listA, numA)
		listB = append(listB, numB)
	}

	return listA, listB
}

func sort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
			}
		}
	}
}

func getDistance(listA []int, listB []int) int {
	if len(listA) != len(listB) {
		msg := fmt.Sprintf("List A does not match List B length: %d != %d", len(listA), len(listB))
		panic(msg)
	}

	distance := 0
	for i := 0; i < len(listA); i++ {
		if listA[i] > listB[i] {
			distance += listA[i] - listB[i]
		} else if listB[i] > listA[i] {
			distance += listB[i] - listA[i]
		}
	}

	return distance
}

func getFrequencyScore(listA []int, listB []int) []int {
	if len(listA) != len(listB) {
		msg := fmt.Sprintf("List A does not match List B length: %d != %d", len(listA), len(listB))
		panic(msg)
	}

	var frequency []int
	for i := 0; i < len(listA); i++ {
		freqCount := 0
		for j := 0; j < len(listA); j++ {
			if listA[i] == listB[j] {
				freqCount++
			}
		}
		frequency = append(frequency, freqCount)
	}

	return frequency
}

func getSimilarityScore(list []int, frequency []int) int {
	similarityScore := 0

	if len(list) != len(frequency) {
		msg := fmt.Sprintf("List length does not match frequency: %d != %d", len(list), len(frequency))
		panic(msg)
	}

	for i := 0; i < len(list); i++ {
		similarityScore += list[i] * frequency[i]
	}

	return similarityScore
}

func main() {
	listA, listB := readFile()
	sort(listA)
	sort(listB)
	distance := getDistance(listA, listB)
	similarityScore := getSimilarityScore(listA, getFrequencyScore(listA, listB))
	fmt.Printf("Total distance: %d\n", distance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
