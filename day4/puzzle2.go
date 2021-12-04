package day4

func puzzle2(boards [][][]int, numbers []int) int {
	locationMap := makeLocationMap(boards)
	bingoBoards := make([]bool, len(boards))
	var winningBoard [][]int
	var winningNumber int
	
	for _, number := range numbers {
		for _, loc := range locationMap[number] {
			if !bingoBoards[loc.board] {
				boards[loc.board][loc.row][loc.col] = -1
				if rowBingo(boards, loc) || colBingo(boards, loc) {
					winningBoard = boards[loc.board]
					winningNumber = number
					bingoBoards[loc.board] = true
				}
			}
		}
		delete(locationMap, number)
	}

	return score(winningBoard, winningNumber)
}
