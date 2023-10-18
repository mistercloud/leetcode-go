package main

// https://leetcode.com/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {

	left := &ListNode{0, nil}
	leftHead := left
	node := head
	var prev *ListNode
	for node != nil {
		if node.Val < x {
			//вставляем в новый список
			left.Next = &ListNode{node.Val, nil}
			left = left.Next
			//удаляем элемент из массива
			if prev == nil {
				head = node.Next
			} else {
				prev.Next = node.Next

			}
		} else {
			prev = node
		}

		node = node.Next
	}
	//объединяем левый с остатками изначального
	left.Next = head
	return leftHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
