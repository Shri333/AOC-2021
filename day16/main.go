package day16

import (
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	data, err := os.ReadFile("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	fmt.Println("Day 16:")
	fmt.Println("Puzzle 1:", puzzle1(input))
	fmt.Println("Puzzle 2:", puzzle2(input))
	fmt.Println()
}
