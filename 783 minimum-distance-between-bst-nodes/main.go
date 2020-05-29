// NOT SOLVED
// NOT SOLVED
// NOT SOLVED

package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	diff, _, _ := getDeeper(root)

	return diff

	/*
		if root == nil {
			return 2147483647
		}

		left, right := 2147483647, 2147483647
		if root.Left != nil {
			leftDiff := root.Val - root.Left.Val
			leftMinDiff := minDiffInBST(root.Left)
			if leftDiff < leftMinDiff {
				left = leftDiff
			} else {
				left = leftMinDiff
			}
		}
		if root.Right != nil {
			rightDiff := root.Right.Val - root.Val
			rightMinDiff := minDiffInBST(root.Right)
			if rightDiff < rightMinDiff {
				right = rightDiff
			} else {
				right = rightMinDiff
			}
		}

		if left < right {
			return left
		} else {
			return right
		}
	*/
}

func getDeeper(root *TreeNode) (int, int, int) {
	if root.Left == nil {
		diffL, minL, maxL := math.MaxInt32, root.Val, root.Val
	} else {
		diffL, minL, maxL := getDeeper(root.Left)
	}
		
	if root.Right == nil {
		diffR, minR, maxR := math.MaxInt32, root.Val, root.Val
	} else {
		diffR, minR, maxR := getDeeper(root.Right)
	}

	diffC := root.Val - maxL
	if diffC > minR - root.Val {
		diffC = minR - root.Val
	}

	if diffL > diffR {
		if diffC > diffR {
			diffC = diffR
		}
	} else if diffC > diffL {
		diffC = diffL
	}
	return diffC, minL, maxR
}
