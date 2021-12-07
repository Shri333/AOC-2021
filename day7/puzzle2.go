package day7

import "math"

func puzzle2(positions []int) int {
	total := math.MaxInt
	lowest, highest := bounds(positions)

	for i := lowest; i <= highest; i++ {
		sum := 0
		for _, position := range positions {
			diff := abs(position - i)
			sum += (diff * (diff + 1)) / 2
		}
		total = min(total, sum)
	}

	return total
}
