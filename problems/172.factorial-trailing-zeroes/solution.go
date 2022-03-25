// 172. 阶乘后的零
// URL：https://leetcode-cn.com/problems/factorial-trailing-zeroes/
// 难度：中等
// 关键词：数学、数论
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 44.92% 的用户

package main

func trailingZeroes(n int) int {
	cnt := 0
	for i := 5; i <= n; i += 5 {
		for j := i; j%5 == 0; j /= 5 {
			cnt++
		}
	}
	return cnt
}
