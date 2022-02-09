// 1447. 最简分数
// URL：https://leetcode-cn.com/problems/simplified-fractions/
// 难度：中等
// 关键词：数论、辗转相除法、欧几里得算法、最大公约数
// 执行用时：48 ms, 在所有 Go 提交中击败了 51.61% 的用户
// 内存消耗：6.9 MB, 在所有 Go 提交中击败了 93.55% 的用户

package main

import "fmt"

func simplifiedFractions(n int) (ans []string) {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	for i := 2; i <= n; i++ {
		ans = append(ans, fmt.Sprintf("1/%v", i))
		for j := 2; j < i; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%v/%v", j, i))
			}
		}
	}
	return
}
