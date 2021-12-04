package day2

import (
	"strconv"
	"strings"
)

func puzzle1(slice []string) int {
	horizontal, depth := 0, 0
	for _, element := range slice {
		movement := strings.Fields(element)
		direction := movement[0]
		change, _ := strconv.Atoi(movement[1])
		switch direction {
		case "forward":
			horizontal += change
		case "down":
			depth += change
		case "up":
			depth -= change
		}
	}

	return horizontal * depth
}
