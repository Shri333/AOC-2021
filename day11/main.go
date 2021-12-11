package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, element := range text {
			num, _ := strconv.Atoi(element)
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	fmt.Println("Day 11:")
	fmt.Println("Puzzle 1:", puzzle1(grid))
	fmt.Println("Puzzle 2:", puzzle2(grid))
	fmt.Println()
}
