package day15

func puzzle2(grid [][]int) int {
	n := len(grid)
	bigGrid := make([][]int, 5*n)
	for i := range bigGrid {
		bigGrid[i] = make([]int, 5*n)
	}

	for I := 0; I < 5*n; I += n {
		for J := 0; J < 5*n; J += n {
			increment := I/n + J/n
			for i := I; i < I+n; i++ {
				for j := J; j < J+n; j++ {
					sum := grid[i-I][j-J] + increment
					if sum > 9 {
						sum -= 9
					}
					bigGrid[i][j] = sum
				}
			}
		}
	}

	return puzzle1(bigGrid)
}
