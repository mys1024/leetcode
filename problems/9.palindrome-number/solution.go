// 9. 回文数
// URL：https://leetcode-cn.com/problems/palindrome-number/
// 难度：简单
// 关键词：数学
// 执行用时：16 ms, 在所有 Go 提交中击败了 64.49% 的用户
// 内存消耗：4.4 MB, 在所有 Go 提交中击败了 98.71% 的用户

package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	reversed := 0
	for tmp != 0 {
		reversed *= 10
		reversed += tmp % 10
		tmp /= 10
	}
	return reversed == x
}
