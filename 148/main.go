package _148

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	right := mid.Next
	mid.Next = nil

	head = sortList(head)
	right = sortList(right)

	return mergeLists(head, right)
}

func mergeLists(head1, head2 *ListNode) *ListNode {
	ret := &ListNode{0, nil}
	head := ret
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head.Next = head1
			head1 = head1.Next
		} else {
			head.Next = head2
			head2 = head2.Next
		}
		head = head.Next
	}

	if head1 != nil {
		head.Next = head1
	}
	if head2 != nil {
		head.Next = head2
	}

	return ret.Next
}

func getMid(head *ListNode) *ListNode {

	slow := head
	fast := head
	var temp *ListNode
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return temp
}

type ListNode struct {
	Val  int
	Next *ListNode
}
