// 5. 最长回文子串
// URL：https://leetcode-cn.com/problems/longest-palindromic-substring/
// 难度：中等
// 关键词：动态规划
// 执行用时：100 ms, 在所有 Go 提交中击败了 35.60% 的用户
// 内存消耗：6.9 MB, 在所有 Go 提交中击败了 36.86% 的用户

package main

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	retStart, retLength := 0, 1
	for end := 1; end < n; end++ {
		for start := 0; start < end; start++ {
			if s[start] != s[end] {
				dp[start][end] = false
			} else {
				if end-start < 3 {
					dp[start][end] = true
				} else {
					dp[start][end] = dp[start+1][end-1]
				}
			}
			length := end - start + 1
			if dp[start][end] && (length > retLength) {
				retStart, retLength = start, length
				if end == n-1 {
					break
				}
			}
		}
	}
	return s[retStart : retStart+retLength]
}
