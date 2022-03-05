// 521. 最长特殊序列 Ⅰ
// URL：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
// 难度：简单
// 关键词：字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 89.80% 的用户

package main

func findLUSlength(a string, b string) int {
	switch {
	case a == b:
		return -1
	case len(a) >= len(b):
		return len(a)
	default:
		return len(b)
	}
}
