package day9

func puzzle1(locations [][]int) int {
	sum := 0
	for i, row := range locations {
		for j, value := range row {
			if i > 0 && value >= locations[i-1][j] {
				continue
			}
			if i < len(locations)-1 && value >= locations[i+1][j] {
				continue
			}
			if j > 0 && value >= locations[i][j-1] {
				continue
			}
			if j < len(row)-1 && value >= locations[i][j+1] {
				continue
			}
			sum += value + 1
		}
	}

	return sum
}
