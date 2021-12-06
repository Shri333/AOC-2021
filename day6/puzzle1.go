package day6

func puzzle1(ages []int, days int) int {
	copy := make([]int, len(ages))
	for i := 0; i < len(ages); i++ {
		copy[i] = ages[i]
	}

	for k := 0; k < days; k++ {
		children := []int{}
		for i := 0; i < len(copy); i++ {
			if copy[i] == 0 {
				copy[i] = 6
				children = append(children, 8)
			} else {
				copy[i]--
			}
		}
		copy = append(copy, children...)
	}

	return len(copy)
}
