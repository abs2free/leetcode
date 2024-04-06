// package main this is a leetcode
package main

/*
*
83. 删除排序链表中的重复元素

给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。



示例 1：
![](assets/list1.jpeg)

输入：head = [1,1,2]
输出：[1,2]
示例 2：
![](assets/list2.jpeg)

输入：head = [1,1,2,3,3]
输出：[1,2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列

*
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	next := slow.Next
	for next != nil {
		if slow.Val == next.Val {
			slow.Next = next.Next
		} else {
			slow = slow.Next
		}

		next = next.Next
	}

	return head
}
