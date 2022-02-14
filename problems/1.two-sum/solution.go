// 1. 两数之和
// URL：https://leetcode-cn.com/problems/two-sum/
// 难度：简单
// 关键词：统计、数组、哈希表
// 执行用时：4 ms, 在所有 Go 提交中击败了 96.57% 的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了 65.04% 的用户

package main

func twoSum(nums []int, target int) []int {
	counter := map[int]int{}
	for i, num := range nums {
		j, ok := counter[target-num]
		if ok {
			return []int{j, i}
		}
		counter[num] = i
	}
	return []int{-1, -1}
}
