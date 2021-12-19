package day19

import "math"

func puzzle2(reports [][]coordinate) int {
	_, scanners := locations(reports)
	max := 0
	for i := 0; i < len(scanners)-1; i++ {
		for j := i + 1; j < len(scanners); j++ {
			s1, s2 := scanners[i], scanners[j]
			dist := int(math.Abs(float64(s1.x-s2.x))) + int(math.Abs(float64(s1.y-s2.y))) + int(math.Abs(float64(s1.z-s2.z)))
			max = int(math.Max(float64(max), float64(dist)))
		}
	}

	return max
}
