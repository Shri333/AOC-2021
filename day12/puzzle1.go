package day12

import (
	"unicode"
)

type set struct {
	data map[string]bool
}

func (s *set) contains(element string) bool {
	_, ok := s.data[element]
	return ok
}

func (s *set) add(element string) bool {
	if s.contains(element) {
		return false
	}

	s.data[element] = true
	return true
}

func (s *set) remove(element string) bool {
	if !s.contains(element) {
		return false
	}

	delete(s.data, element)
	return true
}

func puzzle1(graph map[string][]string) int {
	return dfs1(graph, "start", &set{make(map[string]bool)})
}

func dfs1(graph map[string][]string, current string, visited *set) int {
	if current == "end" {
		return 1
	}

	if isLowerCase(current) {
		visited.add(current)
	}

	paths := 0
	for _, node := range graph[current] {
		if !visited.contains(node) {
			paths += dfs1(graph, node, visited)
		}
	}

	if isLowerCase(current) {
		visited.remove(current)
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
