package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
	z int
}

func RunPuzzles() {
	file, err := os.Open("day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reports := [][]coordinate{}
	current := []coordinate{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			reports = append(reports, current)
			current = []coordinate{}
		} else if !strings.HasPrefix(scanner.Text(), "---") {
			coord := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			z, _ := strconv.Atoi(coord[2])
			current = append(current, coordinate{x, y, z})
		}
	}
	reports = append(reports, current)

	fmt.Println("Day 19:")
	fmt.Println("Puzzle 1:", puzzle1(reports))
	fmt.Println("Puzzle 2:", puzzle2(reports))
	fmt.Println()
}
