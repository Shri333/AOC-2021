package day25

func puzzle1(locations [][]rune) int {
	for step := 1; ; step++ {
		if !move(&locations) {
			return step
		}
	}
}

func move(locations *[][]rune) bool {
	if east(locations) {
		south(locations)
		return true
	}

	return south(locations)
}

func east(locations *[][]rune) bool {
	temp := make([][]rune, len(*locations))
	for i := range temp {
		temp[i] = make([]rune, len((*locations)[0]))
	}

	moved := false
	for i := 0; i < len(*locations); i++ {
		for j := 0; j < len((*locations)[0]); j++ {
			if (*locations)[i][j] != '>' {
				temp[i][j] = (*locations)[i][j]
				continue
			}

			var next int
			if j < len((*locations)[0])-1 {
				next = j + 1
			} else {
				next = 0
			}

			if (*locations)[i][next] == '.' {
				temp[i][j], temp[i][next] = '.', '>'
				moved = true
				j++
			} else {
				temp[i][j] = '>'
			}
		}
	}

	*locations = temp
	return moved
}

func south(locations *[][]rune) bool {
	temp := make([][]rune, len(*locations))
	for i := range temp {
		temp[i] = make([]rune, len((*locations)[0]))
	}

	moved := false
	for j := 0; j < len((*locations)[0]); j++ {
		for i := 0; i < len(*locations); i++ {
			if (*locations)[i][j] != 'v' {
				temp[i][j] = (*locations)[i][j]
				continue
			}

			var next int
			if i < len(*locations)-1 {
				next = i + 1
			} else {
				next = 0
			}

			if (*locations)[next][j] == '.' {
				temp[i][j], temp[next][j] = '.', 'v'
				moved = true
				i++
			} else {
				temp[i][j] = 'v'
			}
		}
	}

	*locations = temp
	return moved
}
