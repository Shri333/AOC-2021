package day10

type runeStack []rune

func (s *runeStack) push(element rune) {
	*s = append(*s, element)
}

func (s *runeStack) pop() rune {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *runeStack) empty() bool {
	return len(*s) == 0
}

func puzzle1(lines []string) int {
	score := 0
	scoreMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	for _, line := range lines {
		s := &runeStack{}
		for _, char := range line {
			if _, ok := charMap[char]; ok {
				s.push(char)
			} else {
				if charMap[s.pop()] != char {
					score += scoreMap[char]
					break
				}
			}
		}
	}

	return score
}
