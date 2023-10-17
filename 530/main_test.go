package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {

	tests := []struct {
		name string
		json string
		want int
	}{
		{
			"Test1",
			"[4,2,6,1,3]",
			1,
		},
		{
			"Test2",
			"[1,0,48,null,null,12,49]",
			1,
		},
		{
			"Test3",
			"[11, null,12 ]",
			1,
		},
		{
			"Test4",
			"[236,104,701,null,227,null,911]",
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := jsonToTree(tt.json)
			Travers(root)
			if got := getMinimumDifference(root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
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
