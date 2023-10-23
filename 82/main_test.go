package _82

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args{[]int{1, 2, 3, 3, 4, 4, 5}},
			[]int{1, 2, 5},
		},
		{
			"Test2",
			args{[]int{1, 1, 1, 2, 3}},
			[]int{2, 3},
		},
		{
			"Test3",
			args{[]int{1, 1, 1, 2, 3, 3, 3}},
			[]int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := deleteDuplicates(arrToList(tt.args.head))
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
