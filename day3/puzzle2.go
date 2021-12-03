package day3

import "strconv"

func puzzle2(data []string) int {
	return getRating(data, true) * getRating(data, false)
}

func getRating(data []string, mostCommon bool) int {
	size := len(data[0])
	matrix := freqMatrix(data, size)

	var rating int64
	set := make(map[string]bool)

	for _, element := range data {
		set[element] = true
	}

	for position := 0; position < size && len(set) > 1; position++ {
		var target byte
		if mostCommon {
			if matrix[position][1] >= matrix[position][0] {
				target = '1'
			} else {
				target = '0'
			}
		} else {
			if matrix[position][0] <= matrix[position][1] {
				target = '0'
			} else {
				target = '1'
			}
		}

		for key := range set {
			if key[position] != target {
				for index, bit := range key {
					matrix[index][int(bit - '0')]--
				}
				delete(set, key)
			}

			if len(set) == 1 {
				break
			}
		}
	}

	for key := range set {
		rating, _ = strconv.ParseInt(key, 2, 64)
	}

	return int(rating)
}
