package day3

import (
	"math"
)

func puzzle1(data []string) int {
	size := len(data[0])
	matrix := freqMatrix(data, size)

	gamma, epsilon := 0, 0
	for i := 0; i < size; i++ {
		if matrix[i][0] > matrix[i][1] {
			epsilon += int(math.Pow(2, float64(size-i-1)))
		} else {
			gamma += int(math.Pow(2, float64(size-i-1)))
		}
	}

	return gamma * epsilon
}

func freqMatrix(data []string, size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = []int{0, 0}
	}

	for _, element := range data {
		for index, bit := range element {
			matrix[index][int(bit-'0')]++
		}
	}

	return matrix
}
