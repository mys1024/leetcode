// 2024. 考试的最大困扰度
// URL：https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
// 难度：中等
// 关键词：滑动窗口
// 执行用时：12 ms, 在所有 Go 提交中击败了 70.97% 的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(str string, k int, c byte) int {
	l, r, ret, n := 0, 0, 0, len(str)
	for r < n {
		if str[r] == c {
			r++
		} else if k > 0 {
			r++
			k--
		} else {
			if str[l] != c {
				k++
			}
			if l == r {
				r++
			}
			l++
		}
		ret = max(ret, r-l)
	}
	return ret
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(count(answerKey, k, 'T'), count(answerKey, k, 'F'))
}
