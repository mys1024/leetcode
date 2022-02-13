// 1324. 竖直打印单词
// URL：https://leetcode-cn.com/problems/print-words-vertically/
// 难度：中等
// 关键词：字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "strings"

func printVertically(s string) []string {
	words := strings.Split(s, " ")
	maxLen := 0
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	ans := []string{}
	for i := 0; i < maxLen; i++ {
		sb := strings.Builder{}
		for _, word := range words {
			if i < len(word) {
				sb.WriteByte(word[i])
			} else {
				sb.WriteByte(' ')
			}
		}
		ans = append(ans, strings.TrimRight(sb.String(), " "))
	}
	return ans
}
