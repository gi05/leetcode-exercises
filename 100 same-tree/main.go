package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []*TreeNode{p, q}
	l := len(stack)

	for l > 0 {
		n1, n2 := stack[l-2], stack[l-1]
		stack = stack[:l-2]

		if n1 == nil && n2 == nil {
			l = len(stack)
			continue
		} else if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		stack = append(stack, n1.Left, n2.Left, n1.Right, n2.Right)
		l = len(stack)
	}
	return true
}
