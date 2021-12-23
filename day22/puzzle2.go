package day22

func (i interval) length() int {
	return i.upper - i.lower + 1
}

func puzzle2(instructions []instruction) int {
	cuboids := []cuboid{}

	count := 0
	for _, c := range cuboids {
		count += c.x.length() * c.y.length() * c.z.length()
	}

	return count
}
