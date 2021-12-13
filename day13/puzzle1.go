package day13

import "fmt"

func puzzle1(paper [][]int, folds []*fold, Y, X int) int {
	fmt.Println(Y, X)
	foldPaper(paper, folds[0], &Y, &X)
	return countDots(paper, Y, X)
}

func foldPaper(paper [][]int, f *fold, Y, X *int) {
	if f.direction == 'x' {
		for i := 0; i <= *Y; i++ {
			for j := 2*f.line - *X; j < f.line; j++ {
				paper[i][j] |= paper[i][2*f.line-j]
			}
		}
		*X = f.line - 1
	} else {
		for i := 2*f.line - *Y; i < f.line; i++ {
			for j := 0; j <= *X; j++ {
				paper[i][j] |= paper[2*f.line-i][j]
			}
		}
		*Y = f.line - 1
	}
}

func countDots(paper [][]int, Y, X int) int {
	count := 0
	for i := 0; i <= Y; i++ {
		for j := 0; j <= X; j++ {
			if paper[i][j] == 1 {
				count++
			}
		}
	}

	return count
}
