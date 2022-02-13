// 540. 有序数组中的单一元素
// URL：https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
// 难度：中等
// 关键词：数组、位运算、二分查找
// 执行用时：20 ms, 在所有 Go 提交中击败了 73.93% 的用户
// 内存消耗：7.8 MB, 在所有 Go 提交中击败了 86.32% 的用户

package main

func singleNonDuplicate(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
