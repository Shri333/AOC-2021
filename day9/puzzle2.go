package day9

import "sort"

func puzzle2(locations [][]int) int {
	sizes := []int{}
	for i := range locations {
		for j := range locations[i] {
			size := dfs(locations, i, j)
			if size != 0 {
				sizes = append(sizes, size)
			}
		}
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	return sizes[0] * sizes[1] * sizes[2]
}

func dfs(locations [][]int, i, j int) int {
	if i < 0 || i >= len(locations) || j < 0 || j >= len(locations[0]) || locations[i][j] == 9 {
		return 0
	}

	locations[i][j] = 9
	return 1 + dfs(locations, i+1, j) + dfs(locations, i-1, j) + dfs(locations, i, j+1) + dfs(locations, i, j-1)
}
