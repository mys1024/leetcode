// 6006. 拿出最少数目的魔法豆
// URL：https://leetcode-cn.com/problems/removing-minimum-number-of-magic-beans/
// 难度：中等
// 关键词：数组、排序
// 执行用时：188 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.5 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import (
	"sort"
)

func minimumRemoval(beans []int) int64 {
	var (
		sum int64 = 0
		min int64 = 999999999999
		n   int   = len(beans)
	)
	for _, bean := range beans {
		sum += int64(bean)
	}
	sort.Ints(beans)
	prev := 0
	for idx, bean := range beans {
		if bean == prev {
			continue
		}
		prev = bean
		r := sum - int64(bean*(n-idx))
		if r < min {
			min = r
		}
	}
	return min
}
