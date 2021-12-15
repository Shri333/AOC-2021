package day13

import "strings"

func puzzle2(paper [][]int, folds []*fold, Y, X int) string {
	for _, f := range folds {
		foldPaper(paper, f, &Y, &X)
	}

	return paperString(paper, Y, X)
}

func paperString(paper [][]int, Y, X int) string {
	sb := &strings.Builder{}
	for i := 0; i <= Y; i++ {
		sb.WriteRune('\n')
		for j := 0; j <= X; j++ {
			if paper[i][j] == 1 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
	}

	return sb.String()
}
