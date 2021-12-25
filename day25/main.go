package day25

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	locations := [][]rune{}
	for scanner.Scan() {
		locations = append(locations, []rune(scanner.Text()))
	}

	fmt.Println("Day 25:")
	fmt.Println("Puzzle 1:", puzzle1(locations))
	fmt.Println("Puzzle 2:", puzzle2(locations))
	fmt.Println()
}
