// MEDIUM
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
func isValidBST(root *TreeNode) bool {
	// if root == nil {
	// 	return true
	// }
	// validL, validR := true, true
	// if root.Left != nil {
	// 	validL = root.Val > root.Left.Val && isValidBST(root.Left)
	// }
	// if root.Right != nil {
	// 	validR = root.Val < root.Right.Val && isValidBST(root.Right)
	// }

	// return validL && validR

	valid, _, _ := helper(root)
	return valid
}

func helper(root *TreeNode) (bool, int, int) {
	// validL, validR, minL, minR, maxL, maxR := true, true, math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	var (
		validL, validR                   bool
		minL, maxL, minR, maxR, min, max int
	)

	if root == nil {
		return true, math.MaxInt32, math.MinInt32
	}
	if root.Left == nil {
		validL, minL, maxL = true, root.Val, root.Val
	} else {
		validL, minL, maxL = helper(root.Left)
		if maxL >= root.Val {
			validL = false
		}
	}
	min = minL

	if root.Right == nil {
		validR, minR, maxR = true, root.Val, root.Val
	} else {
		validR, minR, maxR = helper(root.Right)
		if minR <= root.Val {
			validR = false
		}
	}
	max = maxR

	return validL && validR, min, max
}
