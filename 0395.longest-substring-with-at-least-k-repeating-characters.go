// package main this is a leetcode
package main

/*
*
395. 至少有 K 个重复字符的最长子串

给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

如果不存在这样的子字符串，则返回 0。

示例 1：

输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2：

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

提示：

1 <= s.length <= 104
s 仅由小写英文字母组成
1 <= k <= 105

*
*/
func longestSubstring(s string, k int) int {
	n := len(s)
	if n == 0 || n < k {
		return 0
	}
	if k <= 1 {
		return n
	}

	counts := make(map[byte]int)
	for i := range s {
		char := s[i]
		counts[char]++
	}

	l := 0
	for l < n && counts[s[l]] >= k {
		l++
	}
	if l >= n-1 {
		return l
	}

	ls1 := longestSubstring(s[:l], k)
	// 移除不符合条件的元素
	for l < n && counts[s[l]] < k {
		l++
	}

	var ls2 int
	if l < n {
		ls2 = longestSubstring(s[l:], k)
	} else {
		ls2 = 0
	}

	return max(ls1, ls2)
}
