// 136. 只出现一次的数字
// URL：https://leetcode-cn.com/problems/single-number/
// 难度：简单
// 关键词：位运算
// 执行用时：12 ms, 在所有 Go 提交中击败了 95.95% 的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了 93.47% 的用户

package main

func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
