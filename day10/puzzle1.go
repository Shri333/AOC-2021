package day10

type stack struct {
	data []rune
}

func (s *stack) push(element rune) {
	s.data = append(s.data, element)
}

func (s *stack) pop() rune {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *stack) empty() bool {
	return len(s.data) == 0
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
		s := &stack{[]rune{}}
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
