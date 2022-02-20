// 5997. 找到和为给定整数的三个连续整数
// URL：https://leetcode-cn.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
// 难度：中等
// 关键词：数学
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	a := num / 3
	return []int64{a - 1, a, a + 1}
}
