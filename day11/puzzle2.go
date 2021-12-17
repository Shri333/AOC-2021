package day11

func puzzle2(grid [][]int) int {
	for step := 1; ; step++ {
		increaseEnergyLevel(grid)
		flashes := flashOctupuses(grid)
		if flashes == len(grid)*len(grid[0]) {
			return step
		}

		step++
	}
}
