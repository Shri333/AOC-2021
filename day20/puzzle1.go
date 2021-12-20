package day20

import "math"

func puzzle1(algorithm []int, image [][]int) int {
	enhanced, bit := image, 0
	for i := 0; i < 2; i++ {
		enhanced = enhance(algorithm, enhanced, bit)
		bit = switchBit(algorithm, bit)
	}

	return lit(enhanced)
}

func enhance(algorithm []int, image [][]int, bit int) [][]int {
	padded := pad(image, bit)
	enhanced := make([][]int, len(padded))
	for i := range enhanced {
		enhanced[i] = make([]int, len(padded[0]))
	}

	for i, row := range padded {
		for j := range row {
			enhanced[i][j] = output(algorithm, padded, i, j, bit)
		}
	}

	return enhanced
}

func pad(image [][]int, bit int) [][]int {
	padded := make([][]int, len(image)+2)
	for i := range padded {
		padded[i] = make([]int, len(image[0])+2)
		for j := range padded[i] {
			padded[i][j] = bit
		}
	}

	for i, row := range image {
		for j, value := range row {
			padded[i+1][j+1] = value
		}
	}

	return padded
}

func output(algorithm []int, image [][]int, i, j, bit int) int {
	bits := make([]int, 9)
	if i > 0 {
		if j > 0 {
			bits[0] = image[i-1][j-1]
		} else {
			bits[0] = bit
		}
		bits[1] = image[i-1][j]
		if j < len(image[0])-1 {
			bits[2] = image[i-1][j+1]
		} else {
			bits[2] = bit
		}
	} else {
		bits[0], bits[1], bits[2] = bit, bit, bit
	}

	if j > 0 {
		bits[3] = image[i][j-1]
	} else {
		bits[3] = bit
	}
	bits[4] = image[i][j]
	if j < len(image[0])-1 {
		bits[5] = image[i][j+1]
	} else {
		bits[5] = bit
	}

	if i < len(image)-1 {
		if j > 0 {
			bits[6] = image[i+1][j-1]
		} else {
			bits[6] = bit
		}
		bits[7] = image[i+1][j]
		if j < len(image[0])-1 {
			bits[8] = image[i+1][j+1]
		} else {
			bits[8] = bit
		}
	} else {
		bits[6], bits[7], bits[8] = bit, bit, bit
	}

	return algorithm[decimal(bits)]
}

func decimal(bits []int) int {
	number := 0
	for i, bit := range bits {
		exponent := float64(len(bits) - i - 1)
		number += bit * int(math.Pow(2, exponent))
	}

	return number
}

func switchBit(algorithm []int, bit int) int {
	if bit == 0 && algorithm[0] == 1 {
		bit = 1
	} else if bit == 1 && algorithm[len(algorithm)-1] == 0 {
		bit = 0
	}

	return bit
}

func lit(image [][]int) int {
	count := 0
	for _, row := range image {
		for _, value := range row {
			count += value
		}
	}

	return count
}
