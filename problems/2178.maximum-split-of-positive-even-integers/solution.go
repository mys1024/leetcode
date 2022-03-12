// 2178. 拆分成最多数目的偶整数之和
// URL：https://leetcode-cn.com/problems/maximum-split-of-positive-even-integers/
// 难度：中等
// 关键词：贪心、数学
// 执行用时：60 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.2 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	var (
		nums = []int64{}
		prev int64
	)
	for {
		a := prev + 2
		b := finalSum - a
		if a == b || b <= prev {
			nums = append(nums, finalSum)
			return nums
		}
		nums = append(nums, a)
		prev = a
		finalSum = b
	}
}
