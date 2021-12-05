package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunPuzzles() {
	file, _ := os.Open("day1/input.txt")
	scanner := bufio.NewScanner(file)
	slice := make([]int, 2000)
	current := 0

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		slice[current] = number
		current++
	}

	fmt.Println("Day 1:")
	fmt.Println("Puzzle 1:", puzzle1(slice))
	fmt.Println("Puzzle 2:", puzzle2(slice))
	fmt.Println()
}
