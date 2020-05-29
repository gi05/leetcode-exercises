package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	satisfied := 0

	for kg := len(g) - 1; kg >= 0; kg-- {
		for sk, sv := range s {
			if sv >= g[kg] {
				if sk < len(s)-1 {
					s[sk] = s[len(s)-1]
				}
				s = s[:len(s)-1]
				satisfied++
				break
			}
		}
	}
	return satisfied
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))
}
