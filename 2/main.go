package main

import (
	"fmt"
)

func main() {

	l1, l2 := []int{1}, []int{5, 6, 4}

	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()

	l1, l2 = []int{5, 6, 4}, []int{1}
	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()

	l1, l2 = []int{2, 4, 3}, []int{5, 6, 4}

	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()

	l1, l2 = []int{0}, []int{0}

	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()

	l1, l2 = []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}

	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()

	l1, l2 =
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{5, 6, 4}

	travers(addTwoNumbers(encode(l1), encode(l2)))
	fmt.Println()
}

// https://leetcode.com/problems/add-two-numbers/submissions/1052008065/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1.Val += l2.Val

	if l1.Next == nil && l2.Next != nil {
		l1.Next = addRight(l1, 0)
	}
	if l2.Next == nil && l1.Next != nil {
		l2.Next = addRight(l2, 0)
	}

	if l1.Val >= 10 {
		l1.Val = l1.Val % 10
		if l1.Next == nil {
			l1.Next = addRight(l1, 0)
		}
		l1.Next.Val += 1

	}
	if l1.Next != nil && l2.Next != nil {
		l1.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return l1
}

func travers(list *ListNode) {
	fmt.Print(list.Val)
	if list.Next != nil {
		travers(list.Next)
	}
}

func encode(arr []int) *ListNode {
	root := &ListNode{arr[0], nil}
	next := root
	for i := 1; i < len(arr); i++ {
		next = addRight(next, arr[i])
	}
	return root
}

func addRight(list *ListNode, val int) *ListNode {
	list.Next = &ListNode{val, nil}

	return list.Next

}

type ListNode struct {
	Val  int
	Next *ListNode
}
