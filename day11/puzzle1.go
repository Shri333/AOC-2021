package day11

func puzzle1(grid [][]int) int {
	flashes := 0
	for i := 0; i < 100; i++ {
		increaseEnergyLevel(grid)
		flashes += flashOctupuses(grid)
	}

	return flashes
}

func increaseEnergyLevel(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
		}
	}
}

func flashOctupuses(grid [][]int) int {
	visited := makeVisitedGrid(grid)
	flashes := 0
	flashed := true

	for flashed {
		flashed = false
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] > 9 && !visited[i][j] {
					flashOctupus(grid, visited, i, j)
					flashes++
					flashed = true
				}
			}
		}
	}

	return flashes
}

func makeVisitedGrid(grid [][]int) [][]bool {
	visited := make([][]bool, len(grid))
	for index := range visited {
		visited[index] = make([]bool, len(grid[0]))
	}

	return visited
}

func flashOctupus(grid [][]int, visited [][]bool, i, j int) {
	grid[i][j] = 0
	visited[i][j] = true
	increaseAdjacent(grid, visited, i-1, j-1)
	increaseAdjacent(grid, visited, i, j-1)
	increaseAdjacent(grid, visited, i+1, j-1)
	increaseAdjacent(grid, visited, i-1, j)
	increaseAdjacent(grid, visited, i+1, j)
	increaseAdjacent(grid, visited, i-1, j+1)
	increaseAdjacent(grid, visited, i, j+1)
	increaseAdjacent(grid, visited, i+1, j+1)
}

func increaseAdjacent(grid [][]int, visited [][]bool, i, j int) {
	if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && !visited[i][j] {
		grid[i][j]++
	}
}
