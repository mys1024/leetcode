// 1162. 地图分析
// URL：https://leetcode-cn.com/problems/as-far-from-land-as-possible/
// 难度：中等
// 关键词：广度优先搜索、多源广度优先搜索
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 99.24% 的用户

package main

type Position struct {
	X     int
	Y     int
	Level int
}

func (p0 Position) Add(p Position) Position {
	return Position{p0.X + p.X, p0.Y + p.Y, p0.Level + p.Level}
}

func (p Position) Valid(n int) bool {
	return (0 <= p.X && p.X < n) && (0 <= p.Y && p.Y < n)
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	// 将所有陆地单元格入队
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	queue := make([]Position, 0)
	queueHead := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, Position{i, j, 0})
				vis[i][j] = true
			}
		}
	}
	// 开始 BFS
	maxLevel := 0
	neighborOffsets := []Position{{0, 1, 1}, {0, -1, 1}, {1, 0, 1}, {-1, 0, 1}}
	for queueHead < len(queue) {
		cell := queue[queueHead]
		maxLevel = cell.Level
		queueHead++
		// 将相邻的四个单元格中不超出边界并且尚未访问过的入队
		for _, neighborOffset := range neighborOffsets {
			neighbor := cell.Add(neighborOffset)
			if neighbor.Valid(n) && (!vis[neighbor.X][neighbor.Y]) {
				queue = append(queue, neighbor)
				vis[neighbor.X][neighbor.Y] = true
			}
		}
	}
	// 处理答案
	if maxLevel == 0 {
		return -1
	}
	return maxLevel
}
