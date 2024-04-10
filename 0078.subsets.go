// package main this is a leetcode
package main

import "fmt"

/*
*
78. 子集

*
*/

func subsets(nums []int) [][]int {
	var ans [][]int
	var comb []int

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		fmt.Println(idx, comb)

		// 选
		comb = append(comb, nums[idx])
		dfs(idx + 1)
		// 不选，把上一步加入数组中的移除
		comb = comb[:len(comb)-1]
		dfs(idx + 1)
	}

	dfs(0)
	return ans
}
