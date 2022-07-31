// 2200. 找出数组中的所有 K 近邻下标
// URL：https://leetcode.cn/problems/find-all-k-distant-indices-in-an-array/
// 难度：简单
// 关键词：数组
// 执行用时：68 ms, 在所有 Go 提交中击败了 8.33% 的用户
// 内存消耗：4.7 MB, 在所有 Go 提交中击败了 11.11% 的用户

package main

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	m := map[int]bool{}
	for i := 0; i < n; i++ {
		if nums[i] == key {
			for j := i - k; j <= i+k; j++ {
				if 0 <= j && j < n {
					m[j] = true
				}
			}
		}
	}
	ans := []int{}
	for i := range m {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}
