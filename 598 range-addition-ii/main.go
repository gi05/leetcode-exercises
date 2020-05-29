package main

func maxCount(m int, n int, ops [][]int) int {
	min0, min1 := 40000, 40000
	for _, v := range ops {
		if min0 > v[0] {
			min0 = v[0]
		}
		if min1 > v[1] {
			min1 = v[1]
		}
	}

	if min0 > m {
		min0 = m
	}
	if min1 > n {
		min1 = n
	}
	return min0 * min1
}
