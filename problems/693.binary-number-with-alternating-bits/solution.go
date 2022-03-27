// 693. 交替位二进制数
// URL：https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
// 难度：简单
// 关键词：位运算
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 67.12% 的用户

package main

func hasAlternatingBits(n int) bool {
	flag := n&1 == 1
	n >>= 1
	for n != 0 {
		if flag == (n&1 == 1) {
			return false
		} else {
			flag = !flag
			n >>= 1
		}
	}
	return true
}
