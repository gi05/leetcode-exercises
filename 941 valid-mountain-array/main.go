package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	asc := true
	ascC := 0
	descC := 0

	for k, v := range A {
		if k > 0 {
			if A[k-1] < v {
				if asc {
					ascC++
				} else {
					return false
				}
			} else if A[k-1] > v {
				if asc {
					asc = false
				}
				if !asc {
					descC++
				}
			} else {
				return false
			}
		}
	}
	if ascC > 0 && descC > 0 {
		return true
	}
	return false
}
