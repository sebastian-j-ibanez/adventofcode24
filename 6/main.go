package main

import (
	"bytes"
	"fmt"
	"os"
)

type Symbol byte

const (
	file                  = "input.txt"
	upSymbol       Symbol = '^'
	downSymbol     Symbol = 'V'
	leftSymbol     Symbol = '<'
	rightSymbol    Symbol = '>'
	obstacleSymbol Symbol = '#'
	path           Symbol = 'X'
)

type Direction struct {
	x, y   int
	symbol Symbol
}

type Guard struct {
	x, y int
	dir  Direction
}

type Obstacle struct {
	x, y int
	dir  Direction
}

type Coordinate struct {
	x, y int
	dir  Direction
}

func getGrid() [][]byte {
	var grid [][]byte

	fileContent, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(fileContent, []byte("\n"))
	for _, line := range lines {
		grid = append(grid, line)
	}

	return grid
}

// Part 1 solution
func getGuardPath(grid [][]byte) ([][]byte, int) {
	sum := 0
	up := Direction{0, -1, upSymbol}
	right := Direction{1, 0, rightSymbol}
	down := Direction{0, 1, downSymbol}
	left := Direction{-1, 0, leftSymbol}

	guard := Guard{
		x:   0,
		y:   0,
		dir: up,
	}

	// Find guard
	for i, line := range grid {
		index := bytes.IndexByte(line, byte(upSymbol))
		if index != -1 {
			guard.x = index
			guard.y = i
			break
		}
	}

	maxY := len(grid)
	maxX := len(grid[0])

	for {
		// Check if path has finished
		nextX := guard.x + guard.dir.x
		nextY := guard.y + guard.dir.y
		if nextX > maxX || nextX < 0 ||
			nextY > maxY || nextY < 0 {
			grid[guard.y][guard.x] = byte(path)
			break
		}

		// Process guard's next move
		nextSpot := grid[nextY][nextX]
		if nextSpot == byte(obstacleSymbol) {
			switch guard.dir.symbol {
			case upSymbol:
				guard.dir = right
			case rightSymbol:
				guard.dir = down
			case downSymbol:
				guard.dir = left
			case leftSymbol:
				guard.dir = up
			}
		} else {
			grid[guard.y][guard.x] = byte(path)
			grid[nextY][nextX] = byte(guard.dir.symbol)
			guard.x = nextX
			guard.y = nextY
		}
	}

	for _, line := range grid {
		sum += bytes.Count(line, []byte{byte(path)})
	}

	return grid, sum
}

func guardIsStuck(grid [][]byte, dirHistory [][]Direction, guard Guard) bool {
	up := Direction{0, -1, upSymbol}
	right := Direction{1, 0, rightSymbol}
	down := Direction{0, 1, downSymbol}
	left := Direction{-1, 0, leftSymbol}

	// Find guard
	for i, line := range grid {
		index := bytes.IndexByte(line, byte(upSymbol))
		if index != -1 {
			guard.x = index
			guard.y = i
			break
		}
	}

	maxY := len(grid)
	maxX := len(grid[0])

	for {
		// Check if path has finished
		nextX := guard.x + guard.dir.x
		nextY := guard.y + guard.dir.y
		if nextX >= maxX || nextX < 0 ||
			nextY >= maxY || nextY < 0 {
			grid[guard.y][guard.x] = byte(path)
			break
		}

		// Process guard's next move
		nextSpot := grid[nextY][nextX]
		if nextSpot == byte(obstacleSymbol) {
			switch guard.dir.symbol {
			case upSymbol:
				guard.dir = right
			case rightSymbol:
				guard.dir = down
			case downSymbol:
				guard.dir = left
			case leftSymbol:
				guard.dir = up
			}
		} else if nextSpot == byte(path) && dirHistory[nextY][nextX] == guard.dir {
			return true
		} else {
			grid[guard.y][guard.x] = byte(path)
			grid[nextY][nextX] = byte(guard.dir.symbol)
			guard.x = nextX
			guard.y = nextY
		}
	}

	return false
}

// Part 1 & 2 solution
func getNewObstructionPositions(grid [][]byte) ([][]byte, int, int) {
	dirHistory := make([][]Direction, len(grid))
	for i := range dirHistory {
		dirHistory[i] = make([]Direction, len(grid[i]))
	}
	positions := 0
	sum := 0

	up := Direction{0, -1, upSymbol}
	right := Direction{1, 0, rightSymbol}
	down := Direction{0, 1, downSymbol}
	left := Direction{-1, 0, leftSymbol}

	guard := Guard{
		x:   0,
		y:   0,
		dir: up,
	}

	// Find guard
	for i, line := range grid {
		index := bytes.IndexByte(line, byte(upSymbol))
		if index != -1 {
			guard.x = index
			guard.y = i
			dirHistory[guard.y][guard.x] = up
			break
		}
	}

	guardStartX := guard.x
	guardStartY := guard.y

	// Grid X and Y axis size
	maxY := len(grid)
	maxX := len(grid[0])
	// recentObstacles := []Obstacle{}
	for {
		// Check if guard is inbounds
		// Solution is complete when guard leaves grid
		nextX := guard.x + guard.dir.x
		nextY := guard.y + guard.dir.y
		if nextX >= maxX || nextX < 0 ||
			nextY >= maxY || nextY < 0 {
			grid[guard.y][guard.x] = byte(path)
			break
		}

		// Process guard's next move
		nextSpot := grid[nextY][nextX]
		if nextSpot == byte(obstacleSymbol) {
			// Turn guard to the right
			switch guard.dir.symbol {
			case upSymbol:
				guard.dir = right
			case rightSymbol:
				guard.dir = down
			case downSymbol:
				guard.dir = left
			case leftSymbol:
				guard.dir = up
			}
		} else {
			// Save guard direction at new spot
			dirHistory[nextY][nextX] = guard.dir
			grid[nextY][nextX] = byte(obstacleSymbol)
			stuck := guardIsStuck(grid, dirHistory, Guard{x: guardStartX, y: guardStartY, dir: guard.dir})
			if stuck {
				positions++
			}

			// Move guard forward
			grid[guard.y][guard.x] = byte(path)
			grid[nextY][nextX] = byte(guard.dir.symbol)
			guard.x = nextX
			guard.y = nextY
		}
	}

	// Get number of X's in grid
	for _, line := range grid {
		sum += bytes.Count(line, []byte{byte(path)})
	}

	// for y, line := range grid {
	// 	for x, spot := range line {
	// 		if spot != byte(obstacleSymbol) {
	// 			tmpSymbol := grid[y][x]
	// 			grid[y][x] = byte(obstacleSymbol)
	// 			stuck := guardIsStuck(grid, dirHistory)
	// 			if stuck {
	// 				positions++
	// 			}
	// 			grid[y][x] = tmpSymbol
	// 		}
	// 	}
	// }

	return grid, sum, positions
}

func main() {
	// _, sum := getGuardPath(getGrid())
	_, sum, positions := getNewObstructionPositions(getGrid())
	fmt.Printf("Number of unique spots: %d\n", sum)
	fmt.Printf("Number of potential obstruction positions: %d\n", positions)

}

// // Save the current obstacle
// currObstacle := Obstacle{
// 	x:   guard.x,
// 	y:   guard.y,
// 	dir: guard.dir,
// }

// // Save to recent obstacles (LIFO)
// if len(recentObstacles) >= 3 {
// 	recentObstacles = recentObstacles[:len(recentObstacles)-1]
// 	recentObstacles = append([]Obstacle{currObstacle}, recentObstacles...)
// } else {
// 	recentObstacles = append(recentObstacles, currObstacle)
// }

// // Check for potential obstacle position
// if len(recentObstacles) == 3 {
// 	// Saving recent obstacles in variables for clarity
// 	first := recentObstacles[2]
// 	second := recentObstacles[1]
// 	third := recentObstacles[0]

// 	xMatch := first.x == second.x && second.y == third.y
// 	yMatch := first.y == second.y && second.x == third.x
// 	if xMatch || yMatch {
// 		positions++
// 	}
// }
