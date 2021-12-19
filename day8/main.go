package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	patternsList, digitsList := [][]string{}, [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " | ")
		patternsList = append(patternsList, strings.Split(text[0], " "))
		digitsList = append(digitsList, strings.Split(text[1], " "))
	}

	fmt.Println("Day 8:")
	fmt.Println("Puzzle 1:", puzzle1(digitsList))
	fmt.Println("Puzzle 2:", puzzle2(patternsList, digitsList))
	fmt.Println()
}
