package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice := make([]string, 1000)
	current := 0

	for scanner.Scan() {
		slice[current] = scanner.Text()
		current++
	}

	fmt.Println("Day 2:")
	fmt.Println("Puzzle 1:", puzzle1(slice))
	fmt.Println("Puzzle 2:", puzzle2(slice))
	fmt.Println()
}
