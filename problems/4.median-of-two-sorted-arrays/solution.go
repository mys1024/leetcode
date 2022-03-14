// 4. 寻找两个正序数组的中位数
// URL：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// 难度：困难
// 关键词：数组、并归
// 执行用时：12 ms, 在所有 Go 提交中击败了 82.29% 的用户
// 内存消耗：4.9 MB, 在所有 Go 提交中击败了 77.92% 的用户

// 此题解没有满足题目中所要求的时间复杂度 O(log (m+n))

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	mid := (n + m) / 2
	i, j, k := 0, 0, 0
	a, b := -1, -1
	for k <= mid {
		if i == n {
			a, b = b, nums2[j]
			j++
		} else if j == m {
			a, b = b, nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				a, b = b, nums1[i]
				i++
			} else {
				a, b = b, nums2[j]
				j++
			}
		}
		k++
	}

	if (n+m)%2 == 1 {
		return float64(b)
	}
	return float64(a+b) / 2
}
