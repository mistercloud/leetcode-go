package main

import "fmt"

func main() {
	x := &TreeNode{1, nil, nil}
	x.insertLeft(2)
	x.Left.insertLeft(3)
	x.Left.insertRight(4)
	x.insertRight(5)
	x.Right.insertRight(6)
	flatten(x)
	fmt.Println(x)

	x = nil
	flatten(x)
	fmt.Println(x)
}

var inDepArr []*TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	inDep(root)
	var prev *TreeNode
	for i, node := range inDepArr {
		if i == 0 {
			prev = node
			continue
		}
		prev.Left = nil
		prev.Right = node
		prev = node
	}
}

func inDep(root *TreeNode) {
	if root == nil {
		return
	}
	inDepArr = append(inDepArr, root)
	inDep(root.Left)
	inDep(root.Right)
}

func (node *TreeNode) insertRight(val int) {
	node.Right = &TreeNode{val, nil, nil}
}

func (node *TreeNode) insertLeft(val int) {

	node.Left = &TreeNode{val, nil, nil}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
