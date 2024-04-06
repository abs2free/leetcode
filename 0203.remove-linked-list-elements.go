// package main this is a leetcode
package main

/*
*
203. 移除链表元素

给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。


示例 1：
![](removelinked-list.jpeg)


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
示例 2：

输入：head = [], val = 1
输出：[]
示例 3：

输入：head = [7,7,7,7], val = 7
输出：[]


提示：

列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50

*
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	pre := cur
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		}
		if cur.Val != val {
			pre = cur
		}
		cur = cur.Next
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}
