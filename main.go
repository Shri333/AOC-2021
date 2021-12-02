package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Shri333/AOC-2021/day1"
)

func processData(data []byte) []int {
	slice := make([]int, 2000)
	start, current := 0, 0

	for index, char := range data {
		if char == '\n' {
			str := string(data[start:index])
			num, _ := strconv.Atoi(str)
			slice[current] = num
			start = index + 1
			current++
		}
	}

	return slice
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	slice := processData(data)
	fmt.Println(day1.Puzzle2(slice))
}
