package _148

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test1",
			args{[]int{4, 2, 1, 3}},
			[]int{1, 2, 3, 4},
		},
		{
			"Test2",
			args{[]int{-1, 5, 3, 4, 0}},
			[]int{-1, 0, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sortList(arrToList(tt.args.head))
			got := listToArr(list)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
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
