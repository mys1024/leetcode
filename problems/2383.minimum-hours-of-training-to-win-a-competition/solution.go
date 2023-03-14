// 2383. 赢得比赛需要的最少训练时长
// URL：https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/
// 难度：简单
// 关键词：贪心
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 79.55% 的用户

package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	n, h := len(energy), 0
	eng, exp := initialEnergy, initialExperience
	for i := 0; i < n; i++ {
		exEng, exExp := eng-energy[i], exp-experience[i]
		if exEng <= 0 {
			eng = eng - exEng + 1
			h = h - exEng + 1
		}
		if exExp <= 0 {
			exp = exp - exExp + 1
			h = h - exExp + 1
		}
		eng -= energy[i]
		exp += experience[i]
	}
	return h
}
