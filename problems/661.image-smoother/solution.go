// 661. 图片平滑器
// URL：https://leetcode-cn.com/problems/image-smoother/
// 难度：简单
// 关键词：数组、矩阵
// 执行用时：40 ms, 在所有 Go 提交中击败了 84.44% 的用户
// 内存消耗：7.1 MB, 在所有 Go 提交中击败了 62.22% 的用户

package main

func imageSmoother(img [][]int) [][]int {
	n, m := len(img), len(img[0])

	valid := func(x, y int) bool {
		if x < 0 || n <= x || y < 0 || m <= y {
			return false
		}
		return true
	}

	avg := func(x, y int) int {
		sum, cnt := img[x][y], 1
		offsets := [8][2]int{
			{1, 1}, {1, 0}, {1, -1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, 1}, {0, -1},
		}
		for _, offset := range offsets {
			xx, yy := offset[0]+x, offset[1]+y
			if valid(xx, yy) {
				sum += img[xx][yy]
				cnt++
			}
		}
		return sum / cnt
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ans[i][j] = avg(i, j)
		}
	}

	return ans
}
