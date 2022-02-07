// 883. 三维形体投影面积
// URL：https://leetcode-cn.com/problems/projection-area-of-3d-shapes/
// 难度：简单
// 关键词：立体几何、二维数组
// 执行用时：4 ms, 在所有 Go 提交中击败了 90.91% 的用户
// 内存消耗：3.5 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func projectionArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for x := 0; x < n; x++ {
		yMax := 0
		for y := 0; y < n; y++ {
			count := grid[x][y]
			if grid[x][y] > 0 {
				area++
			}
			if count > yMax {
				yMax = count
			}
		}
		area += yMax
	}
	for y := 0; y < n; y++ {
		xMax := 0
		for x := 0; x < n; x++ {
			count := grid[x][y]
			if count > xMax {
				xMax = count
			}
		}
		area += xMax
	}
	return area
}
