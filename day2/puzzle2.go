package day2

import (
	"strconv"
	"strings"
)

func puzzle2(slice []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, element := range slice {
		movement := strings.Fields(element)
		direction := movement[0]
		change, _ := strconv.Atoi(movement[1])
		switch direction {
		case "down":
			aim += change
		case "up":
			aim -= change
		case "forward":
			horizontal += change
			depth += change * aim
		}
	}

	return horizontal * depth
}
