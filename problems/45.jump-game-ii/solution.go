// 45. 跳跃游戏 II
// URL：https://leetcode-cn.com/problems/jump-game-ii/
// 难度：中等
// 关键词：动态规划
// 执行用时：36 ms, 在所有 Go 提交中击败了 21.42% 的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了 25.67% 的用户

package main

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i, d := range nums {
		for j := i + 1; j < n && j <= i+d; j++ {
			if dp[j] == 0 || (dp[j] != 0 && dp[i]+1 < dp[j]) {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[n-1]
}
