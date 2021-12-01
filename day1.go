package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func puzzle1(data []byte) int {
	prev, start, increases := math.MaxInt32, 0, 0
	for index, element := range data {
		if element == '\n' {
			numstr := string(data[start:index])
			current, _ := strconv.Atoi(numstr)
			if current > prev {
				increases++
			}
			prev = current
			start = index + 1
		}
	}
	return increases
}

func puzzle2(data []byte) int {
	start, windowEnd, increases := 0, 0, 0
	window := make([]int, 3)

	for index, element := range data {
		if element == '\n' {
			numstr := string(data[start:index])
			current, _ := strconv.Atoi(numstr)
			if windowEnd >= 3 {
				oldSum := window[0] + window[1] + window[2]
				newSum := window[1] + window[2] + current
				if newSum > oldSum {
					increases++
				}
				window[0], window[1], window[2] = window[1], window[2], current
			} else {
				window[windowEnd] = current
				windowEnd++
			}
			start = index + 1
		}
	}

	return increases
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(puzzle1(data))
	fmt.Println(puzzle2(data))
}
