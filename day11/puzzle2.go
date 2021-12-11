package day11

func puzzle2(grid [][]int) int {
	for step := 1; ; step++ {
		increaseEnergyLevel(grid)
		visited := makeVisitedGrid(grid)

		flashed := true
		for flashed {
			flashed = false
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] > 9 && !visited[i][j] {
						flashOctupus(grid, visited, i, j)
						flashed = true
					}
				}
			}
		}

		allFlash := true
		for _, row := range visited {
			for _, value := range row {
				if !value {
					allFlash = false
					break
				}
			}
			if !allFlash {
				break
			}
		}

		if allFlash {
			return step + 100
		}
	}
}
