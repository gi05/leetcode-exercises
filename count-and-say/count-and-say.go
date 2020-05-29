package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		var (
			resultString string = ""
			charCount    int    = 0
			prevChar     rune   = 0
		)

		for pos, char := range countAndSay(n - 1) {
			if pos == 0 {
				charCount = 1
				prevChar = char
			} else if char == prevChar {
				charCount++
			} else {
				resultString = fmt.Sprintf("%s%d%c", resultString, charCount, prevChar)
				charCount = 1
				prevChar = char
			}
		}
		resultString = fmt.Sprintf("%s%d%c", resultString, charCount, prevChar)

		return resultString
	}
}

func main() {
	fmt.Println(countAndSay(5))
}
