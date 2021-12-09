package day8

import (
	"sort"
)

func puzzle2(patternsList, digitsList [][]string) int {
	sum := 0

	for i := 0; i < len(patternsList); i++ {
		patterns, digits := patternsList[i], digitsList[i]
		patternMap := processPatterns(patterns)
		sum += getNumber(patternMap, digits)
	}

	return sum
}

func processPatterns(patterns []string) map[string]int {
	for index, pattern := range patterns {
		slice := []rune(pattern)
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
		patterns[index] = string(slice)
	}

	sort.Slice(patterns, func(i, j int) bool { return len(patterns[i]) < len(patterns[j]) })
	patternMap := map[string]int{
		patterns[0]: 1,
		patterns[1]: 7,
		patterns[2]: 4,
		patterns[9]: 8,
	}

	fourOnly := []rune{}
	for _, char := range patterns[2] {
		if byte(char) != patterns[0][0] && byte(char) != patterns[0][1] {
			fourOnly = append(fourOnly, char)
		}
	}

	five := map[string]bool{
		patterns[3]: true,
		patterns[4]: true,
		patterns[5]: true,
	}

	for pattern := range five {
		found1, found4 := 0, 0
		for _, char := range pattern {
			if byte(char) == patterns[0][0] || byte(char) == patterns[0][1] {
				found1++
			}
			if char == fourOnly[0] || char == fourOnly[1] {
				found4++
			}
		}
		if found1 == 2 {
			patternMap[pattern] = 3
			delete(five, pattern)
		}
		if found4 == 2 {
			patternMap[pattern] = 5
			delete(five, pattern)
		}
	}

	for pattern := range five {
		patternMap[pattern] = 2
	}

	six := map[string]bool{
		patterns[6]: true,
		patterns[7]: true,
		patterns[8]: true,
	}

	for pattern := range six {
		found4 := 0
		for _, char := range pattern {
			if char == fourOnly[0] || char == fourOnly[1] {
				found4++
			}
		}
		if found4 == 1 {
			patternMap[pattern] = 0
			delete(six, pattern)
			break
		}
	}

	for pattern := range six {
		found1 := 0
		for _, char := range pattern {
			if byte(char) == patterns[0][0] || byte(char) == patterns[0][1] {
				found1++
			}
		}
		if found1 == 2 {
			patternMap[pattern] = 9
		} else {
			patternMap[pattern] = 6
		}
	}

	return patternMap
}

func getNumber(patternMap map[string]int, digits []string) int {
	for index, digit := range digits {
		slice := []rune(digit)
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
		digits[index] = string(slice)
	}

	number := 0
	for _, digit := range digits {
		number = number*10 + patternMap[digit]
	}

	return number
}
