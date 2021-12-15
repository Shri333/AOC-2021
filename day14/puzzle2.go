package day14

import "math"

func puzzle2(template []rune, rules map[string]rune) int {
	pairs := make(map[string]int)
	for i := 1; i < len(template); i++ {
		key := string(template[i-1 : i+1])
		pairs[key]++
	}

	for step := 0; step < 40; step++ {
		pairs_ := make(map[string]int, len(pairs))
		for pair, freq := range pairs {
			insert := rules[pair]
			pair1 := string([]rune{rune(pair[0]), insert})
			pair2 := string([]rune{insert, rune(pair[1])})
			pairs_[pair1] += freq
			pairs_[pair2] += freq
		}
		pairs = pairs_
	}

	return freqDiff2(pairs)
}

func freqDiff2(pairs map[string]int) int {
	freqMap := make(map[rune][]int)
	for pair, freq := range pairs {
		first, second := rune(pair[0]), rune(pair[1])
		if _, ok := freqMap[first]; ok {
			freqMap[first][0] += freq
		} else {
			freqMap[first] = []int{freq, 0}
		}
		if _, ok := freqMap[second]; ok {
			freqMap[second][1] += freq
		} else {
			freqMap[second] = []int{0, freq}
		}
	}

	min, max := math.MaxInt, 0
	for _, freqs := range freqMap {
		freq := int(math.Max(float64(freqs[0]), float64(freqs[1])))
		if freq > max {
			max = freq
		}
		if freq < min {
			min = freq
		}
	}

	return max - min
}
