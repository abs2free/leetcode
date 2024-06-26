// package main this is a leetcode
package main

/*
*
78. 子集

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的
子集
（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*
*/

func subsets(nums []int) [][]int {
	ans := [][]int{}
	var comb []int

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

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
