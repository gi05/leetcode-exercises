package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	test5 := 0
	output := []bool{}

	for _, v := range A {
		test5 = test5<<1 + v
		test5 %= 5
		if test5 == 0 {
			output = append(output, true)
		} else {
			output = append(output, false)
		}
	}
	return output
}

func main() {
	A := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	fmt.Println(prefixesDivBy5(A))
}
