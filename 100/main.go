package main

// https://leetcode.com/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
