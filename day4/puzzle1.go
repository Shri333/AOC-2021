package day4

type location struct {
	board int
	row   int
	col   int
}

func puzzle1(boards [][][]int, numbers []int) int {
	locationMap := makeLocationMap(boards)
	for _, number := range numbers {
		for _, loc := range locationMap[number] {
			boards[loc.board][loc.row][loc.col] = -1
			if rowBingo(boards, loc) || colBingo(boards, loc) {
				return score(boards[loc.board], number)
			}
		}
		delete(locationMap, number)
	}

	return 0
}

func makeLocationMap(boards [][][]int) map[int][]location {
	locationMap := map[int][]location{}
	for boardIndex, board := range boards {
		for row := range board {
			for col, value := range board[row] {
				locationMap[value] = append(locationMap[value], location{boardIndex, row, col})
			}
		}
	}

	return locationMap
}

func rowBingo(boards [][][]int, loc location) bool {
	board := boards[loc.board]
	for _, value := range board[loc.row] {
		if value != -1 {
			return false
		}
	}

	return true
}

func colBingo(boards [][][]int, loc location) bool {
	board := boards[loc.board]
	for _, row := range board {
		if row[loc.col] != -1 {
			return false
		}
	}

	return true
}

func score(board [][]int, number int) int {
	sum := 0
	for _, row := range board {
		for _, value := range row {
			if value != -1 {
				sum += value
			}
		}
	}

	return sum * number
}
