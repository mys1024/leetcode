// 2016. 增量元素之间的最大差值
// URL：https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements/
// 难度：简单
// 关键词：数组
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "math"

func maximumDifference(nums []int) int {
	ans, min := -1, math.MaxInt
	for _, num := range nums {
		if num > min && num-min > ans {
			ans = num - min
		}
		if num < min {
			min = num
		}
	}
	return ans
}
