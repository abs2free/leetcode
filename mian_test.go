package main

import (
	"golang.org/x/exp/slices"
)

func IntsNotEqual(n1, n2 []int) bool {
	return !equalInts(n1, n2)
}

func equalInts(n1, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}

	slices.Sort(n1)
	slices.Sort(n2)

	for i, n := range n1 {
		if n != n2[i] {
			return false
		}
	}

	return true
}

func LinkedListNotEqual(l1, l2 *ListNode) bool {
	return !equalListNode(l1, l2)
}

func equalListNode(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}
