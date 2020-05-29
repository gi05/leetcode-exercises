package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	weakestM := map[int]int{}
	for kn := 0; kn < len(mat[0]) && len(weakestM) < k; kn++ {
		for km := range mat {
			if mat[km][kn] == 0 {
				if _, ok := weakestM[km]; !ok {
					weakestM[km] = kn
				}
			}
		}
	}

	km := 0
	for len(weakestM) < k {
		if _, ok := weakestM[km]; !ok {
			weakestM[km] = len(mat[0])
		}
		km++
	}

	weakestS := make([]int, len(weakestM))
	wi := 0
	for kw, vw := range weakestM {
		weakestS[wi] = vw<<7 + kw
		wi++
	}
	sort.Ints(weakestS)

	weakestS = weakestS[:k]
	for wk, wv := range weakestS {
		weakestS[wk] = wv & 127
	}

	return weakestS
}
