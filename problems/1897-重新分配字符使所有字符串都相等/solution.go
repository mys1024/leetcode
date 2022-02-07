// 1897. 重新分配字符使所有字符串都相等
// URL：https://leetcode-cn.com/problems/redistribute-characters-to-make-all-strings-equal/
// 难度：简单
// 关键词：字符串、计数
// 执行用时：4 ms, 在所有 Go 提交中击败了 89.19% 的用户
// 内存消耗：3.6 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func makeEqual(words []string) bool {
	n := len(words)
	counter := [26]int{}
	for _, word := range words {
		for _, char := range word {
			counter[int(char-'a')]++
		}
	}
	for _, count := range counter {
		if count%n != 0 {
			return false
		}
	}
	return true
}
