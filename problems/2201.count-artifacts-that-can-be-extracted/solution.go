// 2201. 统计可以提取的工件
// URL：https://leetcode.cn/problems/count-artifacts-that-can-be-extracted/
// 难度：中等
// 关键词：数组、哈希表、模拟
// 执行用时：264 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：28.9 MB, 在所有 Go 提交中击败了 50.00% 的用户

package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	ac := make([]int, len(artifacts))
	for idx, artifact := range artifacts {
		aid := idx + 1
		r1, c1, r2, c2 := artifact[0], artifact[1], artifact[2], artifact[3]
		ac[idx] = (r2 - r1 + 1) * (c2 - c1 + 1)
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				matrix[i][j] = aid
			}
		}
	}

	for _, d := range dig {
		r, c := d[0], d[1]
		aid := matrix[r][c]
		if aid != 0 {
			ac[aid-1]--
		}
	}

	cnt := 0
	for _, c := range ac {
		if c == 0 {
			cnt++
		}
	}
	return cnt
}
