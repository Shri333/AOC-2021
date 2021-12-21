package day21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func RunPuzzles() {
	file, err := os.Open("day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	p1, _ := strconv.Atoi(scanner.Text()[28:])
	scanner.Scan()
	p2, _ := strconv.Atoi(scanner.Text()[28:])

	fmt.Println("Day 21:")
	fmt.Println("Puzzle 1:", puzzle1(p1, p2))
	fmt.Println("Puzzle 2:", puzzle2(p1, p2))
	fmt.Println()
}
