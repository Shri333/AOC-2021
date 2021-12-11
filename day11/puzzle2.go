package day11

func puzzle2(grid [][]int) int {
	step := 1
	for {
		increaseEnergyLevel(grid)
		flashes := flashOctupuses(grid)
		if flashes == len(grid)*len(grid[0]) {
			break
		}

		step++
	}

	return step
}
