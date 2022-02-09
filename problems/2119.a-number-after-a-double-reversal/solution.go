// 2119. 反转两次的数字
// URL：https://leetcode-cn.com/problems/a-number-after-a-double-reversal/
// 难度：简单
// 关键词：数论
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 99.27% 的用户

package main

func isSameAfterReversals(num int) bool {
	return num == 0 || num%10 != 0
}
