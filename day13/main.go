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
	Y, X := 0, 0
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

		Y = int(math.Max(float64(Y), float64(y)))
		X = int(math.Max(float64(X), float64(x)))
	}

	for scanner.Scan() {
		f, text := &fold{}, scanner.Text()
		f.direction = rune(text[11])
		f.line, _ = strconv.Atoi(text[13:])
		folds = append(folds, f)
	}

	fmt.Println("Day 13:")
	fmt.Println("Puzzle 1:", puzzle1(paper, folds, Y, X))
	fmt.Println("Puzzle 2:", puzzle2(paper, folds, Y, X))
	fmt.Println()
}
