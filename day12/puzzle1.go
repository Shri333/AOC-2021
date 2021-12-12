package day12

import "unicode"

func puzzle1(graph map[string][]string) int {
	return dfs1(graph, "start", make(map[string]bool))
}

func dfs1(graph map[string][]string, current string, visited map[string]bool) int {
	if current == "end" {
		return 1
	}

	paths := 0
	for _, node := range graph[current] {
		if node == "start" {
			continue
		}

		if isLowerCase(node) {
			if _, ok := visited[node]; isLowerCase(node) && !ok {
				visited[node] = true
				paths += dfs1(graph, node, visited)
				delete(visited, node)
			}
		} else {
			paths += dfs1(graph, node, visited)
		}
	}

	return paths
}

func isLowerCase(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}

	return true
}
