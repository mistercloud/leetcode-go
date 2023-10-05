package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int

	q := []*TreeNode{root}
	for len(q) > 0 {
		//длинна текущего уровня
		qlen := len(q)
		row := make([]int, qlen)
		qt := []*TreeNode{}
		for i, node := range q {
			row[i] = node.Val
			if node.Left != nil {
				qt = append(qt, node.Left)
			}
			if node.Right != nil {
				qt = append(qt, node.Right)
			}

		}
		result = append(result, row)
		q = qt
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
