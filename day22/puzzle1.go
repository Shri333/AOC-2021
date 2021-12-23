package day22

func puzzle1(instructions []instruction) int {
	grid := make([][][]bool, 101)
	for i := range grid {
		grid[i] = make([][]bool, 101)
		for j := range grid[i] {
			grid[i][j] = make([]bool, 101)
		}
	}

	for _, i := range instructions {
		if i.c.x.lower < -50 || i.c.x.upper > 50 || i.c.y.lower < -50 || i.c.y.upper > 50 || i.c.z.lower < -50 || i.c.z.upper > 50 {
			break
		}

		if i.on {
			for x := i.c.x.lower; x <= i.c.x.upper; x++ {
				for y := i.c.y.lower; y <= i.c.y.upper; y++ {
					for z := i.c.z.lower; z <= i.c.z.upper; z++ {
						grid[50+x][50+y][50+z] = true
					}
				}
			}
		} else {
			for x := i.c.x.lower; x <= i.c.x.upper; x++ {
				for y := i.c.y.lower; y <= i.c.y.upper; y++ {
					for z := i.c.z.lower; z <= i.c.z.upper; z++ {
						grid[50+x][50+y][50+z] = false
					}
				}
			}
		}
	}

	count := 0
	for _, matrix := range grid {
		for _, row := range matrix {
			for _, value := range row {
				if value {
					count++
				}
			}
		}
	}

	return count
}
