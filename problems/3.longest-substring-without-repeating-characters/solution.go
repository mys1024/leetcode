// 3. 无重复字符的最长子串
// URL：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 难度：中等
// 关键词：字符串、滑动窗口、哈希表
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了 96.82% 的用户

package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := [128]int{}
	max := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			next := m[s[j]]
			for i < next {
				m[s[i]] = 0
				i++
			}
			m[s[j]] = j + 1
			if (j - i + 1) > max {
				max = j - i + 1
			}
		}
	}
	return max
}
