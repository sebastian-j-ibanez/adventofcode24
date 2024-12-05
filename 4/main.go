package main

import (
	"bytes"
	"fmt"
	"os"
)

type Direction struct {
	dx, dy int
}

const (
	inputFile = "input.txt"
)

// Part 1

// Get crossword file lines in a 2D slice
func getWordGrid() [][]byte {
	contents, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(contents, []byte{'\n'})
	return lines[:len(lines)-1]
}

// Return true if word is found on grid in specified direction
func findMatchInDirection(grid [][]byte, word []byte, x int, y int, dir Direction) bool {
	xEnd := x + (len(word)-1)*dir.dx
	yEnd := y + (len(word)-1)*dir.dy
	if xEnd < 0 || xEnd >= len(grid) ||
		yEnd < 0 || yEnd >= len(grid) {
		return false
	}

	for j := 1; j < len(word); j++ {
		if grid[y+dir.dy*j][x+dir.dx*j] != word[j] {
			return false
		}
	}

	return true
}

// Search every direction from x,y for possible matches
func findMatches(grid [][]byte, word []byte, x int, y int) int {
	matches := 0

	directions := []Direction{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, d := range directions {
		if findMatchInDirection(grid, word, x, y, d) {
			matches++
		}
	}

	return matches
}

// Get number of word matches
func getXmasWordCount(grid [][]byte, word []byte) int {
	matches := 0
	for y, line := range grid {
		for x, letter := range line {
			if letter == word[0] {
				matches += findMatches(grid, word, x, y)
			}
		}
	}

	return matches
}

// Part 2

// Check if there is a x-mas match in the grid at x,y
func diagonalMatch(grid [][]byte, word []byte, x int, y int, d1 Direction, d2 Direction) bool {
	if len(word) < 3 {
		panic("Word must be at least 3 letters long...")
	}

	x1End := x + d1.dx
	y1End := y + d1.dy
	x2End := x + d2.dx
	y2End := y + d2.dy
	if x1End < 0 || x1End >= len(grid) ||
		y1End < 0 || y1End >= len(grid) ||
		x2End < 0 || x2End >= len(grid) ||
		y2End < 0 || y2End >= len(grid) {
		return false
	}

	gridPoint1 := grid[y+d1.dy][x+d1.dx]
	gridPoint2 := grid[y+d2.dy][x+d2.dx]

	return (gridPoint1 == word[1] && gridPoint2 == word[3]) ||
		(gridPoint1 == word[3] && gridPoint2 == word[1])
}

func checkDiagonalMatch(grid [][]byte, word []byte, x int, y int) bool {
	upRight := Direction{1, 1}
	downLeft := Direction{-1, -1}
	upLeft := Direction{1, -1}
	downRight := Direction{-1, 1}

	return diagonalMatch(grid, word, x, y, upRight, downLeft) &&
		diagonalMatch(grid, word, x, y, upLeft, downRight)
}

// Iterate over grid, searching for 'A'.
func getDiagonalMatchCount(grid [][]byte, word []byte) int {
	matches := 0
	for y, line := range grid {
		for x, letter := range line {
			if letter == word[2] && checkDiagonalMatch(grid, word, x, y) {
				matches++
			}
		}
	}

	return matches
}

func main() {
	wordGrid := getWordGrid()
	word := []byte{'X', 'M', 'A', 'S'}
	count1 := getXmasWordCount(wordGrid, word)
	count2 := getDiagonalMatchCount(wordGrid, word)
	fmt.Printf("Found '%s' %d times!\n", string(word), count1)
	fmt.Printf("Found 'X-%s' %d times!\n", string(word), count2)
}
