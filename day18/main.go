package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []string{}
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	fmt.Println("Day 18:")
	fmt.Println("Puzzle 1:", puzzle1(numbers))
	fmt.Println("Puzzle 2:", puzzle2(numbers))
	fmt.Println()
}
