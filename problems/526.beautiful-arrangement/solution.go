// 526. 优美的排列
// URL：https://leetcode-cn.com/problems/beautiful-arrangement/
// 难度：中等
// 关键词：回溯、状态压缩、动态规划
// 执行用时：48 ms, 在所有 Go 提交中击败了 50.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 98.15% 的用户

package main

func countArrangement(n int) int {
	cnt, idx := 0, 1
	used := make([]bool, n)

	var search func()
	search = func() {
		for i, u := range used {
			if u || ((i+1)%idx != 0 && idx%(i+1) != 0) {
				continue
			}
			if idx != n {
				used[i] = true
				idx++
				search()
				idx--
				used[i] = false
			} else {
				cnt++
			}
		}
	}

	search()
	return cnt
}
