package day8

func puzzle1(digitsList [][]string) int {
	count := 0
	for _, digits := range digitsList {
		for _, digit := range digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}
