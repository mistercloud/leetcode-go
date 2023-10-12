package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name string
		json string
		want []int
	}{
		{
			"Test1",
			"[1,2,3,null,5,null,4]",
			[]int{1, 3, 4},
		},
		{
			"Test2",
			"[1,null,3]",
			[]int{1, 3},
		},
		{
			"Test3",
			"[]",
			[]int{},
		},
		{
			"Test4",
			"[1,2,3,null,5,null,4,1]",
			[]int{1, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := jsonToTree(tt.json)
			if got := rightSideView(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
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
