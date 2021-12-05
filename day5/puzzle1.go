package day5

func puzzle1(segments [][]point) int {
	matrix := make([][]int, 1000)
	for i := range matrix {
		matrix[i] = make([]int, 1000)
	}

	for _, segment := range segments {
		startX, startY := segment[0].x, segment[0].y
		endX, endY := segment[1].x, segment[1].y
		if startX == endX {
			if startY > endY {
				for startY >= endY {
					matrix[startY][startX]++
					startY--
				}
			} else {
				for startY <= endY {
					matrix[startY][startX]++
					startY++
				}
			}
		} else if startY == endY {
			if startX > endX {
				for startX >= endX {
					matrix[startY][startX]++
					startX--
				}
			} else {
				for startX <= endX {
					matrix[startY][startX]++
					startX++
				}
			}
		}
	}

	count := 0
	for _, row := range matrix {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}

	return count
}
