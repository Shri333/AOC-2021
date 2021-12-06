package day6

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzles() {
	bytes, err := os.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(bytes), ",")
	ages := []int{}
	for _, element := range data {
		age, _ := strconv.Atoi(element)
		ages = append(ages, age)
	}

	fmt.Println("Day 6:")
	fmt.Println("Puzzle 1:", puzzle1(ages, 80))
	fmt.Println("Puzzle 2:", puzzle2(ages, 256))
	fmt.Println()
}
