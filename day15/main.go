package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		levels := strings.Split(scanner.Text(), "")
		for _, level := range levels {
			risk, _ := strconv.Atoi(level)
			row = append(row, risk)
		}
		grid = append(grid, row)
	}

	fmt.Println("Day 15:")
	fmt.Println("Puzzle 1:", puzzle1(grid))
	fmt.Println("Puzzle 2:", puzzle2(grid))
	fmt.Println()
}
