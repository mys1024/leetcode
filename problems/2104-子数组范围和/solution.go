// 2104. 子数组范围和
// URL：https://leetcode-cn.com/problems/sum-of-subarray-ranges/
// 难度：中等
// 关键词：动态规划、单调栈
// 执行用时：24 ms, 在所有 Go 提交中击败了 55.67% 的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了 99.51% 的用户

package main

func minAndMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func subArrayRanges(nums []int) (sum int64) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min, max := minAndMax(nums[i], nums[i+1])
		for j := 1; j < n-i; j++ {
			if j > 1 {
				num := nums[i+j]
				if num < min {
					min = num
				} else if num > max {
					max = num
				}
			}
			sum += int64(max - min)
		}
	}
	return
}
