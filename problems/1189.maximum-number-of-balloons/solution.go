// 1189. “气球” 的最大数量
// URL：https://leetcode-cn.com/problems/maximum-number-of-balloons/
// 难度：简单
// 关键词：统计
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 90.63% 的用户

package main

func maxNumberOfBalloons(text string) int {
	counter := [5]int{}
	for _, r := range text {
		switch r {
		case 'b':
			counter[0]++
		case 'a':
			counter[1]++
		case 'l':
			counter[2]++
		case 'o':
			counter[3]++
		case 'n':
			counter[4]++
		}
	}
	counter[2] /= 2
	counter[3] /= 2
	min := 99999999
	for _, count := range counter {
		if count < min {
			min = count
		}
	}
	return min
}
