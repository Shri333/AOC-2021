package day1

func Puzzle2(slice []int) int {
	start, sum, answer := 0, 0, 0
	for index, element := range slice {
		sum += element
		if index >= 3 {
			prev := sum - element
			sum -= slice[start]
			if sum > prev {
				answer++
			}
			start++
		}
	}

	return answer
}
