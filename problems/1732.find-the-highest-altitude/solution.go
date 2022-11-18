// 1725. 可以形成最大正方形的矩形数目
// URL：https://leetcode-cn.com/problems/number-of-rectangles-that-can-form-the-largest-square/
// 难度：简单
// 关键词：数组
// 执行用时：0 ms, 在所有 Python3 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Python3 提交中击败了 97.92% 的用户

package main

func largestAltitude(gain []int) int {
	sum, max := 0, 0
	for _, dh := range gain {
		sum += dh
		if sum > max {
			max = sum
		}
	}
	return max
}
