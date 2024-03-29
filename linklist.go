package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func List2Ints(l *ListNode) (nums []int) {
	tmp := l
	for tmp != nil {
		nums = append(nums, tmp.Val)
		tmp = tmp.Next
	}
	return nums
}

func (ln *ListNode) print() {
	node := ln
	for node != nil {
		fmt.Printf(" ->%d", node.Val)
		node = node.Next
	}
	fmt.Println("")
}

func compareListNode(l1, l2 *ListNode) bool {
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
