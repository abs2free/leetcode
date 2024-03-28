package main

/*
*
234.回文链表

给你一个单链表的头节点 head ，请你判断该链表是否为
回文链表
。如果是，返回 true ；否则，返回 false 。

示例 1：

输入：head = [1,2,2,1]
输出：true
示例 2：

输入：head = [1,2]
输出：false

提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
func isPalindromeLinkedList(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 反转链表
	cloneLinkedList := clone(head)
	prev := reverseList(cloneLinkedList)

	h, t := head, prev
	for h != nil && t != nil {
		if h.Val != t.Val {
			return false
		}

		h = h.Next
		t = t.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func clone(head *ListNode) *ListNode {
	l := &ListNode{}
	t := l
	for head != nil {
		t.Next = &ListNode{Val: head.Val}
		t = t.Next
		head = head.Next
	}
	return l.Next
}
