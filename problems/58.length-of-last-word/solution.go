// 58. 最后一个单词的长度
// URL：https://leetcode-cn.com/problems/length-of-last-word/
// 难度：简单
// 关键词：字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 82.60% 的用户

package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
