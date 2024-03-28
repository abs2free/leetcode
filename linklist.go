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

func (ln *ListNode) print() {
	node := ln
	for node != nil {
		fmt.Printf(" ->%d", node.Val)
		node = node.Next
	}
	fmt.Println("")
}
