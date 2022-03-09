// 79. 单词搜索
// URL：https://leetcode-cn.com/problems/word-search/
// 难度：中等
// 关键词：图、深度优先搜索、回溯算法
// 执行用时：172 ms, 在所有 Go 提交中击败了 10.21% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 99.93% 的用户

package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	isValid := func(x, y int) bool {
		return 0 <= x && x < m && 0 <= y && y < n
	}
	// 原地标记
	mark := func(x, y int) {
		board[x][y] += 'z'
	}
	unmark := func(x, y int) {
		board[x][y] -= 'z'
	}
	isMarked := func(x, y int) bool {
		return board[x][y] > 'z'
	}

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	var search func(x, y, idx int) bool
	search = func(x, y, idx int) bool {
		if !isValid(x, y) || isMarked(x, y) || board[x][y] != word[idx] {
			return false
		} else if idx == len(word)-1 {
			return true
		}
		mark(x, y)
		for _, d := range directions {
			if search(x+d[0], y+d[1], idx+1) {
				return true
			}
		}
		unmark(x, y)
		return false
	}

	for i, row := range board {
		for j := range row {
			if search(i, j, 0) {
				return true
			}
		}
	}
	return false
}
