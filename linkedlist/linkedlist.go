package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ints2List(nums []int) *ListNode {
	ln := &ListNode{}

	t := ln
	for _, num := range nums {
		t.Next = &ListNode{Val: num}
		t = t.Next
	}
	return ln.Next
}

func reverse(head *ListNode) *ListNode {
	var pre, curr *ListNode = nil, head

	for curr != nil {
		nextTmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = nextTmp
	}

	return pre
}

func clone(head *ListNode) *ListNode {
	curr := head

	new := &ListNode{}
	t := new
	for curr != nil {
		t.Next = &ListNode{Val: curr.Val}
		t = t.Next
		curr = curr.Next
	}

	return new.Next
}
