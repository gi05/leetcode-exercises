package main

import "fmt"

func minCostToMoveChips(chips []int) int {
	var odd, even int = 0, 0

	for _, v := range chips {
		if v&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd < even {
		return odd
	}
	return even
}

func main() {
	// chips := []int{2, 2, 2, 3, 3}
	chips := []int{1, 2, 3}
	fmt.Println(minCostToMoveChips(chips))
}
