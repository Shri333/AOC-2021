package day20

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzles() {
	file, err := os.Open("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	algorithm := []int{}
	scanner.Scan()
	for _, r := range scanner.Text() {
		if r == '#' {
			algorithm = append(algorithm, 1)
		} else {
			algorithm = append(algorithm, 0)
		}
	}

	scanner.Scan()

	image := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, r := range scanner.Text() {
			if r == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		image = append(image, row)
	}

	fmt.Println("Day 20:")
	fmt.Println("Puzzle 1:", puzzle1(algorithm, image))
	fmt.Println("Puzzle 2:", puzzle2(algorithm, image))
	fmt.Println()
}
