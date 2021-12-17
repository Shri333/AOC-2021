package day17

import "math"

func puzzle1(x1, x2, y1, y2 int) int {
	max := 0
	for v := 0; v < 1000; v++ {
		y, dy := 0, v
		best := 0
		for y > y2 {
			y += dy
			best = int(math.Max(float64(best), float64(y)))
			dy--
		}

		if y >= y1 {
			max = int(math.Max(float64(max), float64(best)))
		}
	}

	return max
}
