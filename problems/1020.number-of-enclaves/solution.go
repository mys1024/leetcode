// 1020. 飞地的数量
// URL：https://leetcode-cn.com/problems/number-of-enclaves/
// 难度：中等
// 关键词：深度优先搜索、图
// 执行用时：44 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了 98.13% 的用户

package main

func numEnclaves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if n < 3 || m < 3 {
		return 0
	}
	// 深度优先遍历一片连续的陆地，并将这片陆地设为海洋
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || n <= i || j < 0 || m <= j || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}
	// 将与网格边界相连的陆地设为海洋
	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 0; i < m; i++ {
		dfs(0, i)
		dfs(n-1, i)
	}
	// 遍历网格并计算陆地单元格的数量
	cnt := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}
