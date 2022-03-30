// 27. 移除元素
// URL：https://leetcode-cn.com/problems/remove-element/
// 难度：简单
// 关键词：数组、双指针
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 99.97% 的用户

package main

func removeElement(nums []int, val int) int {
	for i, num := range nums {
		if num == val {
			nums[i] = -num - 1
		}
	}
	n, m := len(nums), 0
	for i, j := 0, 0; i < n; i++ {
		if nums[i] >= 0 {
			nums[j] = nums[i]
			j++
			m++
		}
	}
	return m
}
