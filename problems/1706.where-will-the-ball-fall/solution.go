// 1706. 球会落何处
// URL：https://leetcode-cn.com/problems/where-will-the-ball-fall/
// 难度：中等
// 关键词：数组、矩阵。动态规划、深度优先搜索
// 执行用时：20 ms, 在所有 Go 提交中击败了 76.92% 的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func findBall(grid [][]int) []int {
	var (
		m    = len(grid)
		n    = len(grid[0])
		prev = make([]int, n)
		cur  = make([]int, n)
	)
	for i := range prev {
		prev[i] = i
	}
	for i := m - 1; i >= 0; i-- {
		row := grid[i]
		for j, d := range row {
			if (j+d < 0) || (j+d >= n) || d+row[j+d] == 0 {
				cur[j] = -1
			} else {
				cur[j] = prev[j+d]
			}
		}
		copy(prev, cur)
	}
	return prev
}
