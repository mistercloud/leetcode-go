package main

func main() {

}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if right > left {
		return 1 + right
	}

	return 1 + left
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
