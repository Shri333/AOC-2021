package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	fmt.Println("Day 3:")
	fmt.Println("Puzzle 1:", puzzle1(data))
	fmt.Println("Puzzle 2:", puzzle2(data))
}
