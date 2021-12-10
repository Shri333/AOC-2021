package day10

import (
	"sort"
)

func puzzle2(lines []string) int {
	scores := []int{}
	scoreMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	for _, line := range lines {
		s := &stack{[]rune{}}
		corrupted := false
		for _, char := range line {
			if _, ok := charMap[char]; ok {
				s.push(char)
			} else {
				if charMap[s.pop()] != char {
					corrupted = true
					break
				}
			}
		}
		if corrupted {
			continue
		}

		score := 0
		for !s.empty() {
			closing := charMap[s.pop()]
			score = score*5 + scoreMap[closing]
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	return scores[len(scores)/2]
}
