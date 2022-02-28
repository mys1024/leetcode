// 2185. 统计包含给定前缀的字符串
// URL：https://leetcode-cn.com/problems/counting-words-with-a-given-prefix/
// 难度：简单
// 关键词：字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "strings"

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, word := range words {
		if strings.Index(word, pref) == 0 {
			cnt++
		}
	}
	return cnt
}
