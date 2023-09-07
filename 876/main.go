package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	list := initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1, 2, 3, 4, 5, 6}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1, 2, 3, 4, 5, 6, 7}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1, 2, 3, 4, 5, 6, 7, 8}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1, 2}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))

	x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list = initFrom(x)
	fmt.Println(x)
	traverse(list)
	fmt.Println(`result`)
	traverse(middleNode(list))
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

func middleNode(head *ListNode) *ListNode {

	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
