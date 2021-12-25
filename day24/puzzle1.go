package day24

import "strconv"

func puzzle1(instructions []instruction) int {
	largest := 99999999999999
	digits := getDigits(largest)
	for !monad(digits, instructions) {
		largest--
		digits = getDigits(largest)
		for containsZero(digits) {
			largest--
			digits = getDigits(largest)
		}
	}

	return largest
}

func getDigits(model int) []int {
	digits := make([]int, 14)
	current := 13
	for model > 0 {
		digit := model % 10
		digits[current] = digit
		model /= 10
		current--
	}

	return digits
}

func containsZero(digits []int) bool {
	for _, d := range digits {
		if d == 0 {
			return true
		}
	}

	return false
}

func monad(digits []int, instructions []instruction) bool {
	current := 0
	w, x, y, z := 0, 0, 0, 0
	for _, i := range instructions {
		switch i.operation {
		case "inp":
			inp(i, digits, &current, &w, &x, &y, &z)
		case "add":
			add(i, &w, &x, &y, &z)
		case "mul":
			mul(i, &w, &x, &y, &z)
		case "div":
			div(i, &w, &x, &y, &z)
		case "mod":
			mod(i, &w, &x, &y, &z)
		case "eql":
			eql(i, &w, &x, &y, &z)
		}
	}

	return z == 0
}

func inp(i instruction, digits []int, current, w, x, y, z *int) {
	switch i.operand1 {
	case "w":
		*w = digits[*current]
	case "x":
		*x = digits[*current]
	case "y":
		*y = digits[*current]
	case "z":
		*z = digits[*current]
	}
	*current++
}

func add(i instruction, w, x, y, z *int) {
	operand2 := getOperand2(i, w, x, y, z)
	switch i.operand1 {
	case "w":
		*w += operand2
	case "x":
		*x += operand2
	case "y":
		*y += operand2
	case "z":
		*z += operand2
	}
}

func mul(i instruction, w, x, y, z *int) {
	operand2 := getOperand2(i, w, x, y, z)
	switch i.operand1 {
	case "w":
		*w *= operand2
	case "x":
		*x *= operand2
	case "y":
		*y *= operand2
	case "z":
		*z *= operand2
	}
}

func div(i instruction, w, x, y, z *int) {
	operand2 := getOperand2(i, w, x, y, z)
	switch i.operand1 {
	case "w":
		*w /= operand2
	case "x":
		*x /= operand2
	case "y":
		*y /= operand2
	case "z":
		*z /= operand2
	}
}

func mod(i instruction, w, x, y, z *int) {
	operand2 := getOperand2(i, w, x, y, z)
	switch i.operand1 {
	case "w":
		*w %= operand2
	case "x":
		*x %= operand2
	case "y":
		*y %= operand2
	case "z":
		*z %= operand2
	}
}

func eql(i instruction, w, x, y, z *int) {
	operand2 := getOperand2(i, w, x, y, z)
	switch i.operand1 {
	case "w":
		if *w == operand2 {
			*w = 1
		} else {
			*w = 0
		}
	case "x":
		if *x == operand2 {
			*x = 1
		} else {
			*x = 0
		}
	case "y":
		if *y == operand2 {
			*y = 1
		} else {
			*y = 0
		}
	case "z":
		if *z == operand2 {
			*z = 1
		} else {
			*z = 0
		}
	}
}

func getOperand2(i instruction, w, x, y, z *int) int {
	switch i.operand2 {
	case "w":
		return *w
	case "x":
		return *x
	case "y":
		return *y
	case "z":
		return *z
	default:
		operand2, _ := strconv.Atoi(i.operand2)
		return operand2
	}
}
