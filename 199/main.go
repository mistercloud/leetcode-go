package main

// https://leetcode.com/problems/binary-tree-right-side-view/submissions/?envType=study-plan-v2&envId=top-interview-150
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}

	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		qlen := len(q)
		//добавляем первый элемент в массив
		ret = append(ret, q[0].Val)

		for _, node := range q {
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
		}
		q = q[qlen:]
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
