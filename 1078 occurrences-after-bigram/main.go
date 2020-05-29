package main

import (
	"fmt"
	"regexp"
)

func findOcurrences(text string, first string, second string) []string {
	reSpace := regexp.MustCompile(" +")
	reAlpha := regexp.MustCompile("[a-z]+")
	words := reSpace.Split(text, -1)
	if len(words) < 3 {
		return []string{}
	}

	// wordsM := map[string]map[string]bool{}

	occurences := []string{}
	for k, v := range words {
		if k > 0 && k < len(words)-1 {
			if reAlpha.MatchString(v) {
				if words[k-1] == first && v == second {
					occurences = append(occurences, words[k+1])
				}
			}
		}
	}

	return occurences
}

func main() {
	text := "alice is a good girl she is a good student"
	first := "a"
	second := "good"

	fmt.Println(findOcurrences(text, first, second))
}
