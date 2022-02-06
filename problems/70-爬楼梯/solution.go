// 70. 爬楼梯
// URL：https://leetcode-cn.com/problems/climbing-stairs/
// 难度：简单
// 关键词：动态规划、滚动数组
// 执行用时：0 ms, 在所有 Python3 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Python3 提交中击败了 99.24% 的用户

package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	// 滚动数组
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, b+a
	}
	return b
}
