package day20

func puzzle2(algorithm []int, image [][]int) int {
	enhanced, bit := image, 0
	for i := 0; i < 50; i++ {
		enhanced = enhance(algorithm, enhanced, bit)
		bit = switchBit(algorithm, bit)
	}

	return lit(enhanced)
}
