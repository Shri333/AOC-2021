package day12

func puzzle2(graph map[string][]string) int {
	return dfs2(graph, "start", make(map[string]int), true)
}

func dfs2(graph map[string][]string, current string, visited map[string]int, visitable bool) int {
	if current == "end" {
		return 1
	}

	paths := 0
	for _, node := range graph[current] {
		if node == "start" {
			continue
		}

		// credit goes to
		// https://www.reddit.com/r/adventofcode/comments/rehj2r/comment/hob35i3/?utm_source=share&utm_medium=web2x&context=3
		// for helping me greatly simplify backtracking code
		if isLowerCase(node) {
			visited[node]++
			if visited[node] == 1 {
				paths += dfs2(graph, node, visited, visitable)
			} else if visitable && visited[node] == 2 {
				paths += dfs2(graph, node, visited, false)
			}
			visited[node]--
		} else {
			paths += dfs2(graph, node, visited, visitable)
		}
	}

	return paths
}
