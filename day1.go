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
	start, count := 0, 0
	nums := make([]int, 2000)
	for i, b := range data {
		if b == '\n' {
			numstr := string(data[start:i])
			nums[count], _ = strconv.Atoi(numstr)
			start = i + 1
			count++
		}
	}

	start, prev, sum, count := 0, math.MaxInt32, 0, 0
	for i, n := range nums {
		sum += n
		if i >= 2 {
			if sum > prev {
				count++
			}
			prev = sum
			sum -= nums[start]
			start++
		}
	}

	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(puzzle2(data))
}
