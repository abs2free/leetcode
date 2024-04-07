// package main this is a leetcode
package main

import "fmt"

/*
*
5. 最长回文子串

给你一个字符串 s，找到 s 中最长的回文
子串
。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

*
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	l := 0
	ans := ""
	for start := 0; start < n; start++ {
		step := max(l, 1)
		for step <= n {
			end := start + step
			if end > n {
				break
			}
			// 找到了一组回文
			fmt.Println(start, end, ":", s[start:end], isPalindromeSub(s[start:end]), "====")
			if isPalindromeSub(s[start:end]) {
				if end-start > l {
					l = end - start
					ans = s[start:end]
				}
			}
			step++
		}
	}

	return ans
}

func isPalindromeSub(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func longestPalindrome2(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	expand := func(left, right int) (int, int) {
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}

			left--
			right++
		}
		return left + 1, right - 1
	}

	start := 0
	end := 0
	for i := 0; i < n; i++ {
		l1, r1 := expand(i, i)
		l2, r2 := expand(i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}

	}

	return s[start : end+1]
}
