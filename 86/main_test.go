package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test1",
			args{[]int{1, 4, 3, 2, 5, 2}, 3},
			[]int{1, 2, 2, 4, 3, 5},
		},
		{
			"Test2",
			args{[]int{2, 1}, 2},
			[]int{1, 2},
		},
		{
			"Test3",
			args{[]int{}, 2},
			[]int{},
		},
		{
			"Test4",
			args{[]int{1, 2, 3, 4, 5, 6}, 7},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"Test5",
			args{[]int{1, 2, 3, 4, 5, 6}, 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"Test6",
			args{[]int{6, 5, 4, 3, 2, 1}, 2},
			[]int{1, 6, 5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrToList(tt.args.head)
			if got := partition(head, tt.args.x); !reflect.DeepEqual(listToArr(got), tt.want) {
				t.Errorf("partition() = %v, want %v", listToArr(got), tt.want)
			}
		})
	}
}

func arrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNode{arr[0], nil}
	prev := root
	for i := 1; i < len(arr); i++ {
		prev.Next = &ListNode{arr[i], nil}
		prev = prev.Next
	}
	return root
}

func listToArr(node *ListNode) []int {
	ret := []int{}
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}

	return ret
}
