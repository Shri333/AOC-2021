package day21

import "math"

func puzzle1(p1, p2 int) int {
	score1, score2 := 0, 0
	player, count := 1, 0

	for score1 < 1000 && score2 < 1000 {
		first, second, third := checkCount(count+1), checkCount(count+2), checkCount(count+3)
		if player == 1 {
			p1 += first + second + third
			p1 = checkSpace(p1)
			score1 += p1
			player = 2
		} else {
			p2 += first + second + third
			p2 = checkSpace(p2)
			score2 += p2
			player = 1
		}
		count += 3
	}

	return int(math.Min(float64(score1), float64(score2))) * count
}

func checkSpace(p int) int {
	if p > 10 {
		if p%10 == 0 {
			return 10
		}
		return p % 10
	}

	return p
}

func checkCount(count int) int {
	if count > 100 {
		if count%100 == 0 {
			return 100
		}
		return count % 100
	}

	return count
}
