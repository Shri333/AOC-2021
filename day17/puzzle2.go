package day17

func puzzle2(x1, x2, y1, y2 int) int {
	count := 0
	for vx := 1; vx <= x2; vx++ {
		for vy := y1; vy <= 1000; vy++ {
			if shootProbe(vx, vy, x1, x2, y1, y2) {
				count++
			}
		}
	}

	return count
}

func shootProbe(vx, vy, x1, x2, y1, y2 int) bool {
	for x, y, dx, dy := 0, 0, vx, vy; ; {
		x, y = x+dx, y+dy
		if dx > 0 {
			dx--
		}
		dy--

		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			return true
		}

		if x > x2 || y < y1 {
			return false
		}
	}
}
