package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum1, sum2 := 0, 0
	sum := &sum1
	for dk := range distance {
		if dk == start || dk == destination {
			if sum == &sum1 {
				sum = &sum2
			} else {
				sum = &sum1
			}
		}
		*sum += distance[dk]
	}

	if sum1 < sum2 {
		return sum1
	}
	return sum2
}

func main() {
	distance := []int{1, 2, 3, 4}
	start, destination := 2, 0

	fmt.Println(distanceBetweenBusStops(distance, start, destination))
}
