package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	locations := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, location := range strings.Split(scanner.Text(), "") {
			number, _ := strconv.Atoi(location)
			row = append(row, number)
		}
		locations = append(locations, row)
	}

	fmt.Println("Day 9:")
	fmt.Println("Puzzle 1:", puzzle1(locations))
	fmt.Println("Puzzle 2:", puzzle2(locations))
	fmt.Println()
}
