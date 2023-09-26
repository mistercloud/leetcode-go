package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list-ii/
func main() {
	x := []int{1, 2, 3, 4, 5}
	list := initFrom(x)
	left := 2
	right := 4
	traverse(list)
	ret := reverseBetween(list, left, right)
	traverse(ret)

	x = []int{5}
	list = initFrom(x)
	left = 1
	right = 1
	traverse(list)
	ret = reverseBetween(list, left, right)
	traverse(ret)

	x = []int{1, 2}
	list = initFrom(x)
	left = 1
	right = 2
	traverse(list)
	ret = reverseBetween(list, left, right)
	traverse(ret)

	x = []int{1, 2, 3, 4, 5, 6}
	list = initFrom(x)
	left = 2
	right = 5
	traverse(list)
	ret = reverseBetween(list, left, right)
	traverse(ret)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	node := head
	i := 1
	var top *ListNode
	var tail *ListNode
	reverseArr := make([]*ListNode, right-left+1)
	for node != nil && i <= right {
		if i == left-1 {
			top = node
		}
		if i >= left {
			reverseArr[i-left] = node
		}
		node = node.Next
		i++
	}

	tail = node
	for i := len(reverseArr) - 1; i >= 0; i-- {
		if top != nil {
			top.Next = reverseArr[i]
			top = top.Next
		} else {
			head = reverseArr[i]
			top = head
		}
	}

	reverseArr[0].Next = tail
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initFrom(x []int) *ListNode {
	var prev *ListNode
	for i := len(x); i > 0; i-- {
		ele := &ListNode{
			Val: x[i-1],
		}
		if prev != nil {
			ele.Next = prev
		}
		prev = ele
	}
	return prev
}

func traverse(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, ` `)
		head = head.Next
	}
	fmt.Println()
}
