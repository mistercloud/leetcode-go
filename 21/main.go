package main

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	root := &ListNode{}
	next := root

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			next.Next = &ListNode{list1.Val, nil}
			next = next.Next
			list1 = list1.Next
		} else {
			next.Next = &ListNode{list2.Val, nil}
			next = next.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		next.Next = list1
	}
	if list2 != nil {
		next.Next = list2
	}

	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
