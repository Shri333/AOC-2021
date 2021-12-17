package day17

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	data, err := os.ReadFile("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data)[13:], ", ")
	x := strings.Split(input[0][2:], "..")
	y := strings.Split(input[1][2:], "..")
	x1, _ := strconv.Atoi(x[0])
	x2, _ := strconv.Atoi(x[1])
	y1, _ := strconv.Atoi(y[0])
	y2, _ := strconv.Atoi(y[1])

	fmt.Println("Day 17:")
	fmt.Println("Puzzle 1:", puzzle1(x1, x2, y1, y2))
	fmt.Println("Puzzle 2:", puzzle2(x1, x2, y1, y2))
	fmt.Println()
}
