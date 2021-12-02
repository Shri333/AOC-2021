package day1

func Puzzle1(slice []int) int {
	answer := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[i-1] {
			answer++
		}
	}

	return answer
}
