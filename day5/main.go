package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func RunPuzzles() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	segments := [][]point{}

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		start, end := strings.Split(points[0], ","), strings.Split(points[1], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])
		segments = append(segments, []point{{startX, startY}, {endX, endY}})

	}

	fmt.Println("Day 5:")
	fmt.Println("Puzzle 1:", puzzle1(segments))
	fmt.Println("Puzzle 2:", puzzle2(segments))
	fmt.Println()
}
