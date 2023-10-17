package main

import (
	"math"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
func getMinimumDifference(root *TreeNode) int {

	ret := math.MaxInt

	arr := inorderTraversal(root)
	for i := 1; i < len(arr); i++ {
		ret = min(ret, arr[i]-arr[i-1])
	}

	return ret
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
