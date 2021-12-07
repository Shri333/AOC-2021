package day7

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	bytes, err := os.ReadFile("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(bytes), ",")
	positions := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		positions[i], _ = strconv.Atoi(data[i])
	}

	fmt.Println("Day 7:")
	fmt.Println("Puzzle 1:", puzzle1(positions))
	fmt.Println("Puzzle 2:", puzzle2(positions))
	fmt.Println()
}
