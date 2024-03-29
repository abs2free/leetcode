// package main this is a leetcode
package main

/*
*

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := &ListNode{}
	tmp := ll
	incr := 0
	for l1 != nil || l2 != nil || incr != 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + incr
		incr = v / 10
		v = v % 10
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return ll.Next
}
