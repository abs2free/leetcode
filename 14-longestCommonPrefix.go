// package main this is a leetcode
package main

/*
*
14.最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*
*/
func longestCommonPrefix(strs []string) string {
	minLen := 100
	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
		}
	}

	var result string
Loop:
	for i := 0; i < minLen; i++ {
		var prefix string
		for j, str := range strs {
			c := str[i : i+1]
			if j == 0 {
				prefix = c
				continue
			}
			if prefix != c {
				break Loop
			}

		}

		result += prefix
	}

	return result
}

func longestCommonPrefix2(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		strNum := len(strs[i])
		for j := 0; i < len(prefix); j++ {
			if strNum <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}

func longestCommonPrefix3(strs []string) string {
	minLen := 100
	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
		}
	}

	var result []byte
Loop:
	for i := 0; i < minLen; i++ {
		prefix := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if prefix != strs[j][i] {
				break Loop
			}
		}

		result = append(result, prefix)
	}

	return string(result)
}
