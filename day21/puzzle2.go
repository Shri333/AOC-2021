package day21

import "math"

type args struct {
	p1     int
	p2     int
	score1 int
	score2 int
	player int
}

func puzzle2(p1, p2 int) int {
	cache := make(map[args][]int)
	res := play(p1, p2, 0, 0, 1, cache)
	return int(math.Max(float64(res[0]), float64(res[1])))
}

func play(p1, p2, score1, score2, player int, cache map[args][]int) []int {
	if score1 >= 21 {
		return []int{1, 0}
	}

	if score2 >= 21 {
		return []int{0, 1}
	}

	key := args{p1, p2, score1, score2, player}
	if wins, ok := cache[key]; ok {
		return wins
	}

	moves := []int{1, 2, 3}
	cache[key] = []int{0, 0}
	if player == 1 {
		for _, i := range moves {
			for _, j := range moves {
				for _, k := range moves {
					space := checkSpace(p1 + i + j + k)
					res := play(space, p2, score1+space, score2, 2, cache)
					cache[key][0] += res[0]
					cache[key][1] += res[1]
				}
			}
		}
	} else {
		for _, i := range moves {
			for _, j := range moves {
				for _, k := range moves {
					space := checkSpace(p2 + i + j + k)
					res := play(p1, space, score1, score2+space, 1, cache)
					cache[key][0] += res[0]
					cache[key][1] += res[1]
				}
			}
		}
	}

	return cache[key]
}
