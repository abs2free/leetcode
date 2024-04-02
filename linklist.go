package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// newByInts convert []int to List
func newByInts(nums []int) *ListNode {
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

func format2Ints(l *ListNode) (nums []int) {
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
