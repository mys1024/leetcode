// 2379. 得到 K 个黑块的最少涂色次数
// URL：https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// 难度：简单
// 关键词：数组
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 78.00% 的用户

package main

func minimumRecolors(blocks string, k int) int {
	n, cnt := len(blocks), 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			cnt += 1
		}
	}
	min := cnt
	for i := 1; (i + k) <= n; i++ {
		if blocks[i-1] == 'W' {
			cnt -= 1
		}
		if blocks[i+k-1] == 'W' {
			cnt += 1
		}
		if cnt < min {
			min = cnt
		}
	}
	return min
}
