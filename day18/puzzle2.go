package day18

import "math"

func puzzle2(numbers []string) int {
	temp := 0
	ptr := &temp
	max := 0
	for i, x := range numbers {
		for j, y := range numbers {
			if i != j {
				*ptr = 0
				left := parse(x, ptr)
				*ptr = 0
				right := parse(y, ptr)
				sum := &pair{-1, left, right, nil}
				left.parent, right.parent = sum, sum
				sum.reduce()
				max = int(math.Max(float64(max), float64(sum.magnitude())))
			}
		}
	}

	return max
}
