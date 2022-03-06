// 2100. 适合打劫银行的日子
// URL：https://leetcode-cn.com/problems/find-good-days-to-rob-the-bank/
// 难度：中等
// 关键词：数组、动态规划、前缀和
// 执行用时：104 ms, 在所有 Go 提交中击败了 92.62% 的用户
// 内存消耗：8.6 MB, 在所有 Go 提交中击败了 99.18% 的用户

package main

import "math"

func goodDaysToRobBank(security []int, time int) []int {
	goodDays := make([]bool, len(security))
	// 正向遍历 security，将前 time 天非递增的日子标记为 true
	prev, acc := math.MaxInt, -1
	for i, s := range security {
		if s <= prev {
			acc++
		} else {
			acc = 0
		}
		if acc >= time {
			goodDays[i] = true
		}
		prev = s
	}
	// 反向遍历 security，将后 time 天不满足非递减的日子标记为 false
	prev, acc = math.MaxInt, -1
	for i := len(security) - 1; i >= 0; i-- {
		s := security[i]
		if s <= prev {
			acc++
		} else {
			acc = 0
		}
		if goodDays[i] && acc < time {
			goodDays[i] = false
		}
		prev = s
	}
	// 构建答案
	ans := []int{}
	for i, good := range goodDays {
		if good {
			ans = append(ans, i)
		}
	}
	return ans
}
