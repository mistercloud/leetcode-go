package _82

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	ret := &ListNode{0, nil}
	current := ret
	next := head.Next
	for head != nil && head.Next != nil {
		if head.Val != next.Val {
			current.Next = &ListNode{head.Val, nil}
			current = current.Next
			head = next
			next = next.Next
		} else {
			for head != nil {
				if head.Val == next.Val {
					head = head.Next
				} else {
					break
				}
			}
			if head != nil {
				next = head.Next
			}

		}
	}

	if head != nil {
		current.Next = &ListNode{head.Val, nil}
	}

	return ret.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
