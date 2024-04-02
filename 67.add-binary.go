// package main this is a leetcode
package main

/*
*
67. 二进制求和

给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

示例 1：

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"

提示：

1 <= a.length, b.length <= 104
a 和 b 仅由字符 '0' 或 '1' 组成
字符串如果不是 "0" ，就不含前导零

*
*/
func addBinary(a string, b string) string {
	var carry bool
	var res string
	var ans byte
	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		if a[i] == '1' && b[j] == '1' {
			if carry {
				ans = '1'
			} else {
				ans = '0'
			}
			carry = true
		} else if a[i] == '0' && b[j] == '0' {
			if carry {
				ans = '1'
			} else {
				ans = '0'
			}
			carry = false
		} else {
			if carry {
				ans = '0'
				carry = true
			} else {
				ans = '1'
				carry = false
			}
		}

		res = string(ans) + res

		i--
		j--
	}

	for i >= 0 {
		if carry {
			if a[i] == '0' {
				ans = '1'
				carry = false
			} else {
				ans = '0'
				carry = true
			}
		} else {
			ans = a[i]
		}

		res = string(ans) + res
		i--
	}

	for j >= 0 {
		if carry {
			if b[j] == '0' {
				ans = '1'
				carry = false
			} else {
				ans = '0'
				carry = true
			}
		} else {
			ans = a[i]
		}

		res = string(ans) + res
		j--
	}

	if carry {
		res = string('1') + res
	}

	return string(res)
}
