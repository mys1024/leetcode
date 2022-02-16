// 461. 汉明距离
// URL：https://leetcode-cn.com/problems/hamming-distance/
// 难度：简单
// 关键词：位运算
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 87.92% 的用户

package main

func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for i := 0; ; i++ {
		mask := 1 << i
		if mask > z {
			break
		}
		if mask&z > 0 {
			cnt++
		}
	}
	return cnt
}
