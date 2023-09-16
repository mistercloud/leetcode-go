package main

import "fmt"

// https://leetcode.com/problems/binary-tree-inorder-traversal/
func main() {
	x := &TreeNode{1, nil, nil}
	insertRight(x, 2)
	insertLeft(x.Right, 3)
	fmt.Println(x, ` - `, inorderTraversal(x), ` - expected`, []int{1, 3, 2})
	fmt.Println(x, ` - `, inorderTraversal(nil), ` - expected`, []int{})
}

func insertLeft(node *TreeNode, val int) {

	node.Left = &TreeNode{val, nil, nil}
}

func insertRight(node *TreeNode, val int) {
	node.Right = &TreeNode{val, nil, nil}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func arrayToTree(arr []int) *TreeNode {

	return &TreeNode{1, nil, nil}
}
