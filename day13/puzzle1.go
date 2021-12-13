package day13

func puzzle1(paper [][]int, folds []*fold, height, width int) int {
	foldPaper(paper, folds[0], &height, &width)
	return countDots(paper, height, width)
}

func foldPaper(paper [][]int, f *fold, height, width *int) {
	if f.direction == 'x' {
		for i := 0; i < *height; i++ {
			for j := 0; j < f.line; j++ {
				paper[i][j] |= paper[i][*width-j-1]
			}
		}
		*width = f.line
	} else {
		for i := 0; i < f.line; i++ {
			for j := 0; j < *width; j++ {
				paper[i][j] |= paper[*height-i-1][j]
			}
		}
		*height = f.line
	}
}

func countDots(paper [][]int, height, width int) int {
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if paper[i][j] == 1 {
				count++
			}
		}
	}

	return count
}
