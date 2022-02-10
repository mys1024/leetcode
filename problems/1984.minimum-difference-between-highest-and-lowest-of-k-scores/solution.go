// 1984. 学生分数的最小差值
// URL：https://leetcode-cn.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// 难度：简单
// 关键词：排序、滑动窗口
// 执行用时：12 ms, 在所有 Go 提交中击败了 94.74% 的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了 89.47% 的用户

package main

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	min := 999999999999
	for i := 0; i < len(nums)-k+1; i++ {
		d := nums[i+k-1] - nums[i]
		if d < min {
			min = d
		}
	}
	return min
}
