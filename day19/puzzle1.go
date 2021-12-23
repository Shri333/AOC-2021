package day19

func puzzle1(reports [][]coordinate) int {
	beacons, _ := locations(reports)
	return len(beacons)
}

func locations(reports [][]coordinate) (map[coordinate]bool, []coordinate) {
	beacons, scanners := make(map[coordinate]bool), []coordinate{{0, 0, 0}}
	for _, coord := range reports[0] {
		beacons[coord] = true
	}

	initial := make(map[coordinate]bool)
	for _, c := range reports[0] {
		initial[c] = true
	}
	processed := []map[coordinate]bool{initial}
	unprocessed := make([][]coordinate, len(reports)-1)
	copy(unprocessed, reports[1:])
	for len(unprocessed) > 0 {
		for i := len(unprocessed) - 1; i >= 0; i-- {
			found, report1 := false, unprocessed[i]
			for _, report0 := range processed {
				for c1 := range report0 {
					for _, report := range orientations(report1) {
						for _, c2 := range report {
							scanner := coordinate{c1.x - c2.x, c1.y - c2.y, c1.z - c2.z}
							matches := 0
							for _, c := range report {
								coord := coordinate{scanner.x + c.x, scanner.y + c.y, scanner.z + c.z}
								if _, ok := report0[coord]; ok {
									matches++
								}
							}
							if matches >= 12 {
								actual := make(map[coordinate]bool)
								for _, c := range report {
									coord := coordinate{scanner.x + c.x, scanner.y + c.y, scanner.z + c.z}
									actual[coord] = true
									beacons[coord] = true
								}
								scanners = append(scanners, scanner)
								processed = append(processed, actual)
								found = true
								goto FOUND
							}
						}
					}
				}
			}
		FOUND:
			if found {
				unprocessed[i] = unprocessed[len(unprocessed)-1]
				unprocessed = unprocessed[:len(unprocessed)-1]
			}
		}
	}

	return beacons, scanners
}

func orientations(report []coordinate) [][]coordinate {
	// credit goes to https://www.reddit.com/r/adventofcode/comments/rjwhdv/2021_day19_i_need_help_understanding_how_to_solve/
	rotations := []func(c coordinate) coordinate{
		// +x
		func(c coordinate) coordinate { return coordinate{c.x, c.y, c.z} },
		func(c coordinate) coordinate { return coordinate{c.x, -c.z, c.y} },
		func(c coordinate) coordinate { return coordinate{c.x, -c.y, -c.z} },
		func(c coordinate) coordinate { return coordinate{c.x, c.z, -c.y} },
		// +y
		func(c coordinate) coordinate { return coordinate{c.y, c.z, c.x} },
		func(c coordinate) coordinate { return coordinate{c.y, -c.x, c.z} },
		func(c coordinate) coordinate { return coordinate{c.y, -c.z, -c.x} },
		func(c coordinate) coordinate { return coordinate{c.y, c.x, -c.z} },
		// +z
		func(c coordinate) coordinate { return coordinate{c.z, c.x, c.y} },
		func(c coordinate) coordinate { return coordinate{c.z, -c.y, c.x} },
		func(c coordinate) coordinate { return coordinate{c.z, -c.x, -c.y} },
		func(c coordinate) coordinate { return coordinate{c.z, c.y, -c.x} },
		// -x
		func(c coordinate) coordinate { return coordinate{-c.x, c.z, c.y} },
		func(c coordinate) coordinate { return coordinate{-c.x, -c.y, c.z} },
		func(c coordinate) coordinate { return coordinate{-c.x, -c.z, -c.y} },
		func(c coordinate) coordinate { return coordinate{-c.x, c.y, -c.z} },
		// -y
		func(c coordinate) coordinate { return coordinate{-c.y, c.x, c.z} },
		func(c coordinate) coordinate { return coordinate{-c.y, -c.z, c.x} },
		func(c coordinate) coordinate { return coordinate{-c.y, -c.x, -c.z} },
		func(c coordinate) coordinate { return coordinate{-c.y, c.z, -c.x} },
		// -z
		func(c coordinate) coordinate { return coordinate{-c.z, c.y, c.x} },
		func(c coordinate) coordinate { return coordinate{-c.z, -c.x, c.y} },
		func(c coordinate) coordinate { return coordinate{-c.z, -c.y, -c.x} },
		func(c coordinate) coordinate { return coordinate{-c.z, c.x, -c.y} },
	}

	res := [][]coordinate{}
	for _, rotation := range rotations {
		orientation := make([]coordinate, len(report))
		for i, c := range report {
			orientation[i] = rotation(c)
		}
		res = append(res, orientation)
	}

	return res
}
