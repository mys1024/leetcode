// 2038. 如果相邻两个颜色均相同则删除当前颜色
// URL：https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
// 难度：中等
// 关键词：贪心、数学、博弈
// 执行用时：12 ms, 在所有 Go 提交中击败了 90.91% 的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了 72.73% 的用户

package main

func winnerOfGame(colors string) bool {
	n := len(colors)
	cntA, cntB := 0, 0
	for i := 1; i < n-1; i++ {
		x, y, z := colors[i-1], colors[i], colors[i+1]
		if x == 'A' && y == 'A' && z == 'A' {
			cntA++
		}
		if x == 'B' && y == 'B' && z == 'B' {
			cntB++
		}
	}
	return cntA > cntB
}
