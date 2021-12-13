package day13

import (
	"strings"
)

func puzzle2(paper [][]int, folds []*fold, height, width int) string {
	for _, f := range folds {
		foldPaper(paper, f, &height, &width)
	}

	return paperString(paper, height, width)
}

func paperString(paper [][]int, height, width int) string {
	sb := &strings.Builder{}
	for i := 0; i < height; i++ {
		sb.WriteRune('\n')
		for j := 0; j < width; j++ {
			if paper[i][j] == 1 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
	}

	return sb.String()
}
