// package main this is a leetcode
package main

import (
	"slices"
)

/*
*
20.有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

*
*/
func isValidParentheses(s string) bool {
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	st := newStack[rune]()

	for _, c := range s {
		// 如果是左括号，压入栈中
		if slices.Contains([]rune{'{', '[', '('}, c) {
			st.Push(c)
			continue
		}

		// 弹出一个元素
		if st.IsEmpty() {
			return false
		}

		// 匹配元素
		b, ok := m[c]
		if !ok {
			return false
		}
		get := st.Pop()
		if get != b {
			return false
		}
	}

	return st.IsEmpty()
}

type stack[T any] struct {
	values []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

func (s *stack[T]) Pop() T {
	n := len(s.values)
	v := s.values[n-1]
	s.values = s.values[:n-1]
	return v
}
