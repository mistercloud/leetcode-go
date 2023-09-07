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

	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}

	middle := len / 2
	for i := 0; i < middle; i++ {
		head = head.Next
	}
	return head
}
