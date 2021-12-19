package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunPuzzles() {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(map[string][]string)
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	fmt.Println("Day 12:")
	fmt.Println("Puzzle 1:", puzzle1(graph))
	fmt.Println("Puzzle 2:", puzzle2(graph))
	fmt.Println()
}
