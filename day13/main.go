package day13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	direction rune
	line      int
}

func RunPuzzles() {
	paper := make([][]int, 2000)
	for index := range paper {
		paper[index] = make([]int, 2000)
	}
	height, width := 0, 0
	folds := []*fold{}

	file, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		location := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(location[0])
		y, _ := strconv.Atoi(location[1])
		paper[y][x] = 1

		height = int(math.Max(float64(height), float64(y)))
		width = int(math.Max(float64(width), float64(x)))
	}

	height++
	width++

	if height%2 == 0 {
		height++
	}

	if width%2 == 0 {
		width++
	}

	for scanner.Scan() {
		f, text := &fold{}, scanner.Text()
		f.direction = rune(text[11])
		f.line, _ = strconv.Atoi(text[13:])
		folds = append(folds, f)
	}

	fmt.Println("Day 13:")
	fmt.Println("Puzzle 1:", puzzle1(paper, folds, height, width))
	fmt.Println("Puzzle 2:", puzzle2(paper, folds, height, width))
	fmt.Println()
}
