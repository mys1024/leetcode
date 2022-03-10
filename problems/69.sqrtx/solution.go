// 69. x 的平方根
// URL：https://leetcode-cn.com/problems/sqrtx/
// 难度：简单
// 关键词：数学、二分搜索
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 81.61% 的用户

package main

func mySqrt(x int) int {
	min := 1
	max := x
	mid := (min + max) / 2
	for !(mid*mid <= x && (mid+1)*(mid+1) > x) {
		if mid*mid < x {
			min = mid
		} else {
			max = mid
		}
		mid = (min + max) / 2
	}
	return mid
}
