package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var leftS, rightS []int = []int{}, []int{}

	leftS = flatten(root.Left, true)
	rightS = flatten(root.Right, false)

	if len(leftS) != len(rightS) {
		return false
	}

	for k, _ := range leftS {
		if leftS[k] != rightS[k] {
			return false
		}
	}

	return true
}

func flatten(root *TreeNode, leftFirst bool) []int {
	if root == nil {
		return []int{-1}
	}

	var leftS, rightS []int
	leftS = flatten(root.Left, leftFirst)
	rightS = flatten(root.Right, leftFirst)

	if leftFirst {
		return append(append([]int{root.Val}, leftS...), rightS...)
	} else {
		return append(append([]int{root.Val}, rightS...), leftS...)
	}
}
