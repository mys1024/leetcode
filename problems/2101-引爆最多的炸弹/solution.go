// 2101. 引爆最多的炸弹
// URL：https://leetcode-cn.com/problems/detonate-the-maximum-bombs/
// 难度：中等
// 关键词：图、有向图、深度优先搜索
// 执行用时：72 ms, 在所有 Go 提交中击败了 32.35% 的用户
// 内存消耗：5.2 MB, 在所有 Go 提交中击败了 94.12% 的用户

package main

import (
	"math"
)

func count(index int, edges [][]bool, marks []bool) int {
	// 对有向图进行深度优先遍历并计数
	counter := 1
	marks[index] = true
	for i, edge := range edges[index] {
		if edge && (!marks[i]) {
			marks[i] = true
			counter += count(i, edges, marks)
		}
	}
	return counter
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	// 创建存储有向图的边的二维数组
	edges := make([][]bool, n)
	for i := range edges {
		edges[i] = make([]bool, n)
	}
	// 扫描有向图的边
	for i, bomb0 := range bombs {
		x0, y0, r0 := bomb0[0], bomb0[1], bomb0[2]
		for j, bomb1 := range bombs {
			if i == j {
				edges[i][j] = true
				continue
			}
			x1, y1 := bomb1[0], bomb1[1]
			edges[i][j] = math.Pow(float64(r0), 2) >= math.Pow(float64(x0-x1), 2)+math.Pow(float64(y0-y1), 2)
		}
	}
	// 遍历每个炸弹作为有向图的起点，计算若该炸弹引爆
	// 总共能同时引爆多少个炸弹，记录最大值并返回。
	max := 0
	for i := range edges {
		func() {
			c := count(i, edges, make([]bool, n))
			if c > max {
				max = c
			}
		}()
	}
	return max
}
