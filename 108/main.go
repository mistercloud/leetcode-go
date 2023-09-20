package main

import "fmt"

func main() {

	nums := []int{0, 1, 2, 3, 4, 5}
	tree := sortedArrayToBST(nums)
	fmt.Println(nums, tree)
	Travers(tree)
	fmt.Println()

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	//выбираем рутом
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func Travers(tree *TreeNode) {
	if tree.Left != nil {
		Travers(tree.Left)
	}
	fmt.Print(tree.Val, ` `)
	if tree.Right != nil {
		Travers(tree.Right)
	}

}
