package tree_node

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func jsonToTree(data string) *TreeNode {
	var intslice []int
	json.Unmarshal([]byte(data), &intslice)

	if len(intslice) == 0 {
		return nil
	}
	if len(intslice) == 1 {
		return &TreeNode{intslice[0], nil, nil}
	}

	root := &TreeNode{intslice[0], nil, nil}
	intslice = intslice[1:]

	q := []*TreeNode{root}
	for len(intslice) > 0 {
		node := q[0]
		q = q[1:]
		val := intslice[0]
		intslice = intslice[1:]
		if val != 0 {
			node.Left = &TreeNode{val, nil, nil}
			q = append(q, node.Left)
		}

		if len(intslice) > 0 {
			val = intslice[0]
			intslice = intslice[1:]
			if val != 0 {
				node.Right = &TreeNode{val, nil, nil}
				q = append(q, node.Right)
			}
		}
	}

	fmt.Println(intslice)

	return root
}

func Travers(tree *TreeNode) {
	if tree != nil {
		if tree.Left != nil {
			Travers(tree.Left)
		}
		fmt.Print(tree.Val, ` `)
		if tree.Right != nil {
			Travers(tree.Right)
		}
	}
}
