// package main this is a leetcode
package main

/*
*
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长
子串

	的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

*
*/
func lengthOfLongestSubstring(s string) int {
	l := 0
	m := make(map[byte]int)
	i := 0
	for j := 0; j < len(s); j++ {
		c := s[j]
		i = max(i, m[c])
		m[c] = j + 1
		l = max(l, j-i+1)
	}
	return l
}

func lengthOfLongestSubstring2(s string) int {
	longestSubLen := 0
	charMap := make(map[byte]bool)

	start := 0
	for end := range s {
		char := s[end]

		if !charMap[char] {
			charMap[char] = true
			longestSubLen = max(longestSubLen, end-start+1)
		} else {
			for charMap[char] {
				charMap[s[start]] = false
				start++
			}
			charMap[char] = true
		}
	}

	return longestSubLen
}
