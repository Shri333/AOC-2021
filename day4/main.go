package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitWhitespace(text string) []string {
	data := []string{}
	current := ""
	for _, char := range text {
		if char == ' ' {
			if current != "" {
				data = append(data, current)
			}
			current = ""
		} else {
			current += string(char)
		}
	}

	data = append(data, current)
	return data
}

func RunPuzzles() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	if scanner.Scan() {
		numstrs := strings.Split(scanner.Text(), ",")
		for _, numstr := range numstrs {
			number, _ := strconv.Atoi(numstr)
			numbers = append(numbers, number)
		}
		scanner.Scan()
	}

	boards := [][][]int{}
	board := make([][]int, 5)
	row := 0
	for scanner.Scan() {
		numstrs := splitWhitespace(scanner.Text())
		for _, numstr := range numstrs {
			number, _ := strconv.Atoi(numstr)
			board[row] = append(board[row], number)
		}

		if row == 4 {
			scanner.Scan()
			boards = append(boards, board)
			board = make([][]int, 5)
			row = 0
		} else {
			row++
		}
	}

	fmt.Println("Day 4:")
	fmt.Println("Puzzle 1:", puzzle1(boards, numbers))
	fmt.Println("Puzzle 2:", puzzle2(boards, numbers))
	fmt.Println()
}
