package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		json string
		want [][]int
	}{
		{
			`Test 1`,
			`[3,9,20,null,null,15,7]`,
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			`Test 2`,
			`[1]`,
			[][]int{{1}},
		},
		{
			`Test 3`,
			`[]`,
			[][]int{},
		},
		{
			`Test 4`,
			`[3,9,20,1,2,15,7]`,
			[][]int{{3}, {9, 20}, {1, 2, 15, 7}},
		},
		{
			`Test 4`,
			`[3,9,20,1,null,15,7]`,
			[][]int{{3}, {9, 20}, {1, 15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//формируем дерево на основе файла json
			node := jsonToTree(tt.json)
			if got := levelOrder(node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
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
