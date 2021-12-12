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

		if isLowerCase(node) {
			visited[node]++
			if visitable && visited[node] == 2 {
				paths += dfs2(graph, node, visited, false)
			} else if visited[node] < 2 {
				paths += dfs2(graph, node, visited, visitable)
			}
			visited[node]--
		} else {
			paths += dfs2(graph, node, visited, visitable)
		}
	}

	return paths
}
