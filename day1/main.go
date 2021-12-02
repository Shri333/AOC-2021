package day1

import (
	"fmt"
	"os"
	"strconv"
)

func RunPuzzles() {
	data, _ := os.ReadFile("input.txt")
	slice := make([]int, 2000)
	start, current := 0, 0

	for index, char := range data {
		if char == '\n' {
			str := string(data[start:index])
			num, _ := strconv.Atoi(str)
			slice[current] = num
			start = index + 1
			current++
		}
	}

	fmt.Println("Day 1:")
	fmt.Println("Puzzle 1: ", puzzle1(slice))
	fmt.Println("Puzzle 2: ", puzzle2(slice))
	fmt.Println()
}
