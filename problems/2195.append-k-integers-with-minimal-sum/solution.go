// 2195. 向数组中追加 K 个整数
// URL：https://leetcode-cn.com/problems/append-k-integers-with-minimal-sum/
// 难度：中等
// 关键词：数组、贪心
// 执行用时：120 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：9.4 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "sort"

func sn(a0, n int) int64 {
	return int64(n)*int64(a0) + int64(n)*int64(n-1)/2
}

func minimalKSum(nums []int, k int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sum := int64(0)
	prev := 0
	for _, num := range nums {
		vacancy := num - prev - 1
		if vacancy > 0 {
			if vacancy > k {
				return sum + sn(prev+1, k)
			} else {
				sum += sn(prev+1, vacancy)
				k -= (num - prev - 1)
			}
		}
		prev = num
	}
	return sum + sn(nums[len(nums)-1]+1, k)
}
