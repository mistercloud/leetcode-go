package main

import "fmt"

func main() {

	x := &TreeNode{0, nil, nil}
	fmt.Println(evaluateTree(x))

	x = &TreeNode{2, nil, nil}
	x.Left = &TreeNode{1, nil, nil}
	x.Right = &TreeNode{3, nil, nil}
	x.Right.Left = &TreeNode{0, nil, nil}
	x.Right.Right = &TreeNode{1, nil, nil}
	fmt.Println(evaluateTree(x))
}

// https://leetcode.com/problems/evaluate-boolean-binary-tree/submissions/1051679295/
func evaluateTree(root *TreeNode) bool {

	if root.Val == 0 {
		return false
	}
	if root.Val == 1 {
		return true
	}

	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	if root.Val == 3 {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
