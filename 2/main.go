package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input = "input.txt"
)

// PART 1
func getRooms() [][]int {
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	roomStrings := strings.Split(string(content), "\n")
	var rooms [][]int
	for _, roomString := range roomStrings {
		if roomString == "" {
			continue
		}
		roomNumbers := strings.Split(roomString, " ")

		var room []int
		for i := range roomNumbers {
			number, err := strconv.Atoi(roomNumbers[i])
			if err != nil {
				panic(err)
			}
			room = append(room, number)
		}

		rooms = append(rooms, room)
	}

	return rooms
}

func getSafeRoomCount(rooms [][]int) int {
	safeRoomCount := 0

	for _, room := range rooms {
		if roomIsSafe(room) {
			safeRoomCount++
		}
	}

	return safeRoomCount
}

func roomIsSafe(room []int) bool {
	initialDirection := room[1] - room[0]
	if initialDirection == 0 || initialDirection < -3 || initialDirection > 3 {
		return false
	}

	for i := 1; i < len(room); i++ {
		diff := room[i] - room[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		if (initialDirection > 0 && diff < 0) || (initialDirection < 0 && diff > 0) {
			return false
		}
	}

	return true
}

// PART 2
func getDampenerSafeRoomCount(rooms [][]int) int {
	safeRoomCount := 0

	for _, room := range rooms {
		if canBeMadeSafe(room) {
			safeRoomCount++
		}
	}

	return safeRoomCount
}

func canBeMadeSafe(room []int) bool {
	for i := 0; i < len(room); i++ {
		if !roomIsSafe(room) {
			tmpRoom := append([]int{}, room...)
			tmpRoom = append(tmpRoom[:i], tmpRoom[i+1:]...)
			if roomIsSafe(tmpRoom) {
				return true
			}
		}
	}

	return false
}

func main() {
	rooms := getRooms()
	safeRoomCount := getSafeRoomCount(rooms)
	dampenerSafeRoomCount := getDampenerSafeRoomCount(rooms)
	fmt.Printf("Safe rooms: %d\n", safeRoomCount)
	fmt.Printf("Safe rooms with dampener: %d\n", dampenerSafeRoomCount)
	fmt.Printf("Total: %d\n", safeRoomCount+dampenerSafeRoomCount)
}
