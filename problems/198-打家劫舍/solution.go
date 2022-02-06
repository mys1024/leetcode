// 198. 打家劫舍
// URL：https://leetcode-cn.com/problems/house-robber/
// 难度：中等
// 关键词：动态规划、滚动数组
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 99.51% 的用户

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	prev, cur := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		prev, cur = cur, max(cur, prev+nums[i])
	}
	return cur
}
