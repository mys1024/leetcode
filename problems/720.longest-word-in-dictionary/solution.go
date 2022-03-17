// 720. 词典中最长的单词
// URL：https://leetcode-cn.com/problems/longest-word-in-dictionary/
// 难度：简单
// 关键词：字典树、哈希表、字符串
// 执行用时：12 ms, 在所有 Go 提交中击败了 83.48% 的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了 93.91% 的用户

package main

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)

	m := map[string]int{}
	for i, word := range words {
		m[word] = i
	}

	valid := make([]bool, len(words))
	for i, word := range words {
		if len(word) == 1 {
			valid[i] = true
		} else {
			idx, ok := m[word[:len(word)-1]]
			if !ok {
				valid[i] = false
			} else {
				valid[i] = valid[idx]
			}
		}
	}

	longest := ""
	for i, word := range words {
		if valid[i] && len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}
