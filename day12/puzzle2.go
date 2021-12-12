package day12

type twice struct {
	node string
	freq int
}

func puzzle2(graph map[string][]string) int {
	paths, unique := &set{make(map[string]bool)}, &set{make(map[string]bool)}

	for node := range graph {
		if node != "start" && isLowerCase(node) {
			unique.add(node)
		}
	}

	for node := range unique.data {
		dfs2(graph, "start", &set{make(map[string]bool)}, &twice{node, 0}, paths, "")
	}

	return len(paths.data)
}

func dfs2(graph map[string][]string, current string, visited *set, t *twice, paths *set, path string) {
	path += current

	if current == "end" {
		paths.add(path)
		return
	}

	if isLowerCase(current) {
		if current == t.node {
			t.freq++
		} else {
			visited.add(current)
		}
	}

	for _, node := range graph[current] {
		if node == t.node && t.freq == 2 {
			continue
		}

		if !visited.contains(node) {
			dfs2(graph, node, visited, t, paths, path)
		}
	}

	if isLowerCase(current) {
		if current == t.node {
			t.freq--
		} else {
			visited.remove(current)
		}
	}
}
