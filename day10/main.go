package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Day 10:")
	fmt.Println("Puzzle 1:", puzzle1(lines))
	fmt.Println("Puzzle 2:", puzzle2(lines))
	fmt.Println()
}
