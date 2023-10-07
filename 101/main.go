package main

// https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {

	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
		qlen := len(q)
		//проверяем симетричен ли массив
		if len(q)%2 == 1 {
			return false
		}
		half := qlen / 2
		for i := 0; i < half; i++ {
			if q[i] == nil && q[len(q)-i-1] != nil {
				return false
			}
			if q[i] != nil && q[len(q)-i-1] == nil {
				return false
			}
			if q[i] != nil && q[len(q)-i-1] != nil && q[i].Val != q[len(q)-i-1].Val {
				return false
			}
		}
		for _, node := range q {
			if node != nil {
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
		q = q[qlen:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
