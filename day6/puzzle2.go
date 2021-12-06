package day6

func puzzle2(ages []int, days int) int {
	count := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, age := range ages {
		count[age]++
	}

	for i := 0; i < days; i++ {
		temp := count[0]
		for j := 0; j < 8; j++ {
			count[j] = count[j+1]
		}
		count[6] += temp
		count[8] = temp
	}

	sum := 0
	for _, element := range count {
		sum += element
	}

	return sum
}
