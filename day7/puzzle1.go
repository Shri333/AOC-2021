package day7

import (
	"math"
)

func puzzle1(positions []int) int {
	total := math.MaxInt
	lowest, highest := bounds(positions)

	for i := lowest; i <= highest; i++ {
		sum := 0
		for _, position := range positions {
			sum += abs(position - i)
		}
		total = min(total, sum)
	}

	return total
}

func bounds(positions []int) (int, int) {
	lowest, highest := math.MaxInt, math.MinInt
	for _, position := range positions {
		lowest = min(lowest, position)
		highest = max(highest, position)
	}

	return lowest, highest
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
