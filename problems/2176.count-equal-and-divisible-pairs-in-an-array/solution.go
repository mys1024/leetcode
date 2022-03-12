// 2176. 统计数组中相等且可以被整除的数对
// URL：https://leetcode-cn.com/problems/count-equal-and-divisible-pairs-in-an-array/
// 难度：简单
// 关键词：统计
// 执行用时：4 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func countPairs(nums []int, k int) int {
	cnt := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] == nums[j]) && ((i*j)%k == 0) {
				cnt++
			}
		}
	}
	return cnt
}
