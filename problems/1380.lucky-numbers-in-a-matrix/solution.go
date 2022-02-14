// 1380. 矩阵中的幸运数
// URL：https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
// 难度：简单
// 关键词：矩阵
// 执行用时：16 ms, 在所有 Go 提交中击败了 88.64% 的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了 97.73% 的用户

package main

func luckyNumbers(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	rowMin := make([]int, n)
	colMax := make([]int, m)
	for i := 0; i < n; i++ {
		min := 99999999
		for j := 0; j < m; j++ {
			num := matrix[i][j]
			if num < min {
				min = num
			}
		}
		rowMin[i] = min
	}
	for j := 0; j < m; j++ {
		max := 0
		for i := 0; i < n; i++ {
			num := matrix[i][j]
			if num > max {
				max = num
			}
		}
		colMax[j] = max
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := matrix[i][j]
			if num == rowMin[i] && num == colMax[j] {
				ans = append(ans, num)
			}
		}
	}
	return ans
}
