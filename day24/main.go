package day24

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type instruction struct {
	operation string
	operand1  string
	operand2  string
}

func RunPuzzles() {
	file, err := os.Open("day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		if len(data) == 2 {
			instructions = append(instructions, instruction{data[0], data[1], ""})
		} else {
			instructions = append(instructions, instruction{data[0], data[1], data[2]})
		}
	}

	fmt.Println("Day 24:")
	fmt.Println("Puzzle 1:", puzzle1(instructions))
	fmt.Println("Puzzle 2:", puzzle2(instructions))
	fmt.Println()
}
