// 2207. 字符串中最多数目的子字符串
// URL：https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string/
// 难度：中等
// 关键词：字符串、前缀和
// 执行用时：16 ms, 在所有 Go 提交中击败了 68.42% 的用户
// 内存消耗：8.2 MB, 在所有 Go 提交中击败了 5.26% 的用户

package main

func maximumSubsequenceCount(text string, pattern string) int64 {
	a, b := pattern[0], pattern[1]
	if a == b {
		aCnt, cnt := 0, int64(0)
		for _, r := range text {
			if byte(r) == a {
				cnt += int64(aCnt)
				aCnt++
			}
		}
		return cnt + int64(aCnt)
	} else {
		aCnt, bCnt, cnt := 0, 0, int64(0)
		pre := make([]int, len(text))
		for i, c := range text {
			if byte(c) == a {
				aCnt++
			}
			pre[i] = aCnt
		}
		for i, c := range text {
			if byte(c) == b {
				bCnt++
				cnt += int64(pre[i])
			}
		}
		if aCnt > bCnt {
			cnt += int64(aCnt)
		} else {
			cnt += int64(bCnt)
		}
		return cnt
	}
}
