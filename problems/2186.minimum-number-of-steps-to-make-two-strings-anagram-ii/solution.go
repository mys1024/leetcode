// 2186. 使两字符串互为字母异位词的最少步骤数
// URL：https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
// 难度：中等
// 关键词：字符串、统计、哈希表
// 执行用时：124 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：7 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func minSteps(s string, t string) int {
	runes1 := map[rune]int{}
	runes2 := map[rune]int{}
	cnt := 0
	for _, r := range s {
		runes1[r]++
	}
	for _, r := range t {
		runes2[r]++
	}
	for r, c := range runes1 {
		if c < runes2[r] {
			cnt += runes2[r] - c
			runes1[r] = runes2[r]
		} else if c > runes2[r] {
			cnt += c - runes2[r]
			runes2[r] = c
		}
	}
	for r, c := range runes2 {
		if c < runes1[r] {
			cnt += runes1[r] - c
			runes2[r] = runes1[r]
		} else if c > runes1[r] {
			cnt += c - runes1[r]
			runes1[r] = c
		}
	}
	return cnt
}
