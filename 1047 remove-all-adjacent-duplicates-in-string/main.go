package main

import "fmt"

func removeDuplicates(S string) string {
	if len(S) < 2 {
		return S
	}

	stack := []byte{}

	for k := 0; k < len(S); k++ {
		if len(stack) == 0 || stack[len(stack)-1] != S[k] {
			stack = append(stack, S[k])
		} else {
			if stack[len(stack)-1] == S[k] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)
}

func main() {
	S := "cabbaca"

	fmt.Println(removeDuplicates(S))
}
