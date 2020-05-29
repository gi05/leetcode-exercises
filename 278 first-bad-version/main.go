package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		if isBadVersion(1) {
			return 1
		}
		return 0
	}

	lBound, rBound, pointer := 1, n, n/2
	bad1, bad2 := true, true
	for {
		bad1 = isBadVersion(pointer)
		bad2 = isBadVersion(pointer + 1)
		if !bad1 && bad2 {
			return pointer + 1
		} else if bad1 && bad2 {
			rBound = pointer
		} else {
			lBound = pointer + 1
		}
		pointer = lBound + (rBound-lBound)/2

		if lBound == rBound {
			return lBound
		}
	}

	return 0
}
