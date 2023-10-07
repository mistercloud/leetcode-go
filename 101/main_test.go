package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			`Test1`,
			"[1,2,2,3,4,4,3]",
			true,
		},
		{
			`Test2`,
			"[1,2,2,null,3,null,3]",
			false,
		},
		{
			`Test3`,
			"[1]",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := jsonToTree(tt.data)
			if got := isSymmetric(node); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
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
