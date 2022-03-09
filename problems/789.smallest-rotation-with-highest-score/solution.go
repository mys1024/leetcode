// 798. 得分最高的最小轮调
// URL：https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/
// 难度：困难
// 关键词：数组、差分、前缀和
// 执行用时：88 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：9.3 MB, 在所有 Go 提交中击败了 66.67% 的用户

package main

func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n+1) // 差分数组，长度 +1 避免越界访问

	for i, num := range nums {
		if i >= num {
			diffs[0]++
			diffs[i-num+1]--
			diffs[i+1]++
			diffs[n]--
		} else {
			diffs[i+1]++
			diffs[i+n-num+1]--
		}
	}

	score, maxScore, maxScoreIdx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore = score
			maxScoreIdx = i
		}
	}
	return maxScoreIdx
}
