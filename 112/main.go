package main

func main() {

}

// https://leetcode.com/problems/path-sum/description/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return root.Left == nil && root.Right == nil && root.Val == targetSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
