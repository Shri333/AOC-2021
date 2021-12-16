package day15

import (
	"container/heap"
	"math"
)

type cell struct {
	i     int
	j     int
	dist  int
	index int
}

func (c *cell) neighbors(visited [][]bool) []int {
	n, adjacent := len(visited), []int{}
	if c.i > 0 && !visited[c.i-1][c.j] {
		adjacent = append(adjacent, n*(c.i-1)+c.j)
	}
	if c.i < n-1 && !visited[c.i+1][c.j] {
		adjacent = append(adjacent, n*(c.i+1)+c.j)
	}
	if c.j > 0 && !visited[c.i][c.j-1] {
		adjacent = append(adjacent, n*c.i+(c.j-1))
	}
	if c.j < n-1 && !visited[c.i][c.j+1] {
		adjacent = append(adjacent, n*c.i+(c.j+1))
	}

	return adjacent
}

type gridPQ []*cell

func (pq *gridPQ) Len() int           { return len(*pq) }
func (pq *gridPQ) Less(i, j int) bool { return (*pq)[i].dist < (*pq)[j].dist }
func (pq *gridPQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *gridPQ) Push(c interface{}) { *pq = append(*pq, c.(*cell)) }
func (pq *gridPQ) Pop() interface{} {
	c := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return c
}

func puzzle1(grid [][]int) int {
	cache := make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid))
	}

	return dijkstra(grid)
}

// credit goes to: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func dijkstra(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid))
	}

	pq := &gridPQ{}
	dists := make([][]int, n)
	for i := range dists {
		dists[i] = make([]int, n)
		for j := range dists[i] {
			if i == 0 && j == 0 {
				dists[i][j] = 0
			} else {
				dists[i][j] = math.MaxInt16
			}
			heap.Push(pq, &cell{i, j, dists[i][j], pq.Len()})
		}
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*cell)
		for _, neighbor := range current.neighbors(visited) {
			i, j := neighbor/n, neighbor%n
			tentative := dists[current.i][current.j] + grid[i][j]
			if tentative < dists[i][j] {
				dists[i][j] = tentative
				var update *cell
				for _, c := range *pq {
					if c.i == i && c.j == j {
						update = c
						break
					}
				}
				update.dist = tentative
				heap.Fix(pq, update.index)
			}
			visited[i][j] = true
		}
	}

	return dists[n-1][n-1]
}
