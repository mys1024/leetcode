// 1922. 统计好数字的数目
// URL：https://leetcode-cn.com/problems/count-good-numbers/
// 难度：中等
// 关键词：快速幂、数学
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

// 带模数的快速幂
func fastPow(x, y, m int64) int64 {
	if y == 0 {
		return 1
	}
	x = x % m
	if y == 1 {
		return x
	}
	z := int64(1)
	for y > 1 {
		if y%2 == 1 {
			z = (x * z) % m
		}
		x = (x * x) % m
		y /= 2
	}
	return (x * z) % m
}

func countGoodNumbers(n int64) int {
	m := int64(1e9 + 7)
	cnt := fastPow(20, n/2, m)
	if n%2 == 1 {
		cnt = (cnt * 5) % m
	}
	return int(cnt)
}
