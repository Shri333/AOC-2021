package day16

import "math"

var binaryMap = map[rune][]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

type packet interface{}

type literal struct {
	version int
	value   int
}

type operator struct {
	version int
	typeID  int
	packets []packet
}

func puzzle1(input string) int {
	bits := getBits(input)
	temp := 0
	ptr := &temp
	return versionSum(parsePacket(bits, ptr))
}

func getBits(input string) []int {
	bits := make([]int, 4*len(input))
	for i, r := range input {
		binary := binaryMap[r]
		for j := 4 * i; j < 4*(i+1); j++ {
			bits[j] = binary[j-4*i]
		}
	}

	return bits
}

func parsePacket(bits []int, i *int) packet {
	version := decimal(bits[*i : *i+3])
	*i += 3

	typeID := decimal(bits[*i : *i+3])
	*i += 3

	if typeID == 4 {
		return parseLiteral(bits, i, version)
	}

	return parseOperator(bits, i, version, typeID)
}

func parseLiteral(bits []int, i *int, version int) *literal {
	binary := []int{}
	current := bits[*i : *i+5]
	for current[0] != 0 {
		binary = append(binary, current[1:]...)
		*i += 5
		current = bits[*i : *i+5]
	}
	binary = append(binary, current[1:]...)
	*i += 5

	return &literal{version, decimal(binary)}
}

func parseOperator(bits []int, i *int, version, typeID int) *operator {
	op := &operator{version, typeID, []packet{}}

	lengthTypeID := bits[*i]
	*i++

	if lengthTypeID == 0 {
		length := decimal(bits[*i : *i+15])
		*i += 15
		temp := *i
		for *i < temp+length {
			op.packets = append(op.packets, parsePacket(bits, i))
		}
	} else {
		length := decimal(bits[*i : *i+11])
		*i += 11
		for j := 0; j < length; j++ {
			op.packets = append(op.packets, parsePacket(bits, i))
		}
	}

	return op
}

func decimal(bits []int) int {
	number := 0
	for i := 0; i < len(bits); i++ {
		exponent := len(bits) - i - 1
		number += bits[i] * int(math.Pow(2, float64(exponent)))
	}

	return number
}

func versionSum(p packet) int {
	switch node := p.(type) {
	case *literal:
		return node.version
	case *operator:
		sum := node.version
		for _, child := range node.packets {
			sum += versionSum(child)
		}
		return sum
	}

	return 0
}
