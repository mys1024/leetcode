// 2044. 统计按位或能得到最大值的子集数目
// URL：https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
// 难度：中等
// 关键词：数组、回溯
// 执行用时：4 ms, 在所有 Go 提交中击败了 76.32% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 94.74% 的用户

package main

func countMaxOrSubsets(nums []int) int {
	n, max, cnt := len(nums), 0, 0

	var dfs func(idx, val int)
	dfs = func(idx, val int) {
		for i := idx; i < n; i++ {
			newVal := val | nums[i]
			if newVal == max {
				cnt++
			} else if newVal > max {
				max = newVal
				cnt = 1
			}
			dfs(i+1, newVal)
		}
	}
	dfs(0, 0)

	return cnt
}
