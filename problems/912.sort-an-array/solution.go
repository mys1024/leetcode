// 912. 排序数组
// URL：https://leetcode-cn.com/problems/sort-an-array/
// 难度：中等
// 关键词：数组、排序、快速排序
// 执行用时：48 ms, 在所有 Go 提交中击败了 72.92% 的用户
// 内存消耗：7.4 MB, 在所有 Go 提交中击败了 67.11% 的用户

package main

import "math/rand"

func QuickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	randIdx := rand.Intn(n)
	pivot := nums[randIdx]
	nums[randIdx] = nums[0]
	i, j := 0, n-1
	for i != j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	QuickSort(nums[:i])
	QuickSort(nums[i+1:])
}

func sortArray(nums []int) []int {
	QuickSort(nums)
	return nums
}
