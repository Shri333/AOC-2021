package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := []rune(scanner.Text())

	scanner.Scan()

	rules := make(map[string]rune)
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), " -> ")
		rules[rule[0]] = []rune(rule[1])[0]
	}

	template_ := make([]rune, len(template))
	copy(template_, template)

	fmt.Println("Day 14:")
	fmt.Println("Puzzle 1:", puzzle1(template, rules))
	fmt.Println("Puzzle 2:", puzzle2(template_, rules))
	fmt.Println()
}
