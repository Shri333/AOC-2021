package day16

import "math"

func puzzle2(input string) int {
	bits := getBits(input)
	temp := 0
	i := &temp
	return eval(parsePacket(bits, i))
}

func eval(p packet) int {
	switch node := p.(type) {
	case *literal:
		return node.value
	case *operator:
		switch node.typeID {
		case 0:
			sum := 0
			for _, child := range node.packets {
				sum += eval(child)
			}
			return sum
		case 1:
			product := 1
			for _, child := range node.packets {
				product *= eval(child)
			}
			return product
		case 2:
			minimum := math.MaxInt32
			for _, child := range node.packets {
				minimum = int(math.Min(float64(minimum), float64(eval(child))))
			}
			return minimum
		case 3:
			maximum := 0
			for _, child := range node.packets {
				maximum = int(math.Max(float64(maximum), float64(eval(child))))
			}
			return maximum
		case 5:
			if eval(node.packets[0]) > eval(node.packets[1]) {
				return 1
			} else {
				return 0
			}
		case 6:
			if eval(node.packets[0]) < eval(node.packets[1]) {
				return 1
			} else {
				return 0
			}
		case 7:
			if eval(node.packets[0]) == eval(node.packets[1]) {
				return 1
			} else {
				return 0
			}
		}
	}

	return 0
}
