package main

import "fmt"

func longestPalindrome(s string) int {
	count := map[rune]int{}

	for _, c := range s {

		if _, ok := count[c]; ok {
			count[c]++
		} else {
			count[c] = 1
		}
	}

	odd := 0
	evenLen := 0
	for _, v := range count {
		if v > 1 {
			evenLen += v - v&1
		}
		if v&1 == 1 {
			odd = 1
		}
	}

	return evenLen + odd
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
