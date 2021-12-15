package day14

import "math"

func puzzle1(template []rune, rules map[string]rune) int {
	for step := 0; step < 10; step++ {
		for i := 1; i < len(template); i++ {
			pair := string(template[i-1 : i+1])
			template = append(template[:i+1], template[i:]...)
			template[i] = rules[pair]
			i++
		}
	}

	return freqDiff1(template)
}

func freqDiff1(template []rune) int {
	freqMap := make(map[rune]int)
	for _, r := range template {
		freqMap[r]++
	}

	min, max := math.MaxInt, 0
	for _, freq := range freqMap {
		if freq > max {
			max = freq
		}
		if freq < min {
			min = freq
		}
	}

	return max - min
}
