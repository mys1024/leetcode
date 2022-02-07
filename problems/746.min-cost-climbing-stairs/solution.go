// 746. 使用最小花费爬楼梯
// URL：https://leetcode-cn.com/problems/min-cost-climbing-stairs/
// 难度：简单
// 关键词：动态规划、滚动数组
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了 100.00% 的用户

// 本题与 "70.爬楼梯" 相似

package main

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		a, b = b, min(a, b)+cost[i]
	}
	return min(a, b)
}
