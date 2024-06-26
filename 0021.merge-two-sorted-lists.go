// package main this is a leetcode
package main

/*
*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
![](assets/merge_ex1.jpeg)

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

*
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ll := &ListNode{}
	t := ll
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			t.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			t.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		t = t.Next
	}

	for list1 != nil {
		t.Next = &ListNode{Val: list1.Val}
		t = t.Next
		list1 = list1.Next
	}

	for list2 != nil {
		t.Next = &ListNode{Val: list2.Val}
		t = t.Next
		list2 = list2.Next
	}
	return ll.Next
}
