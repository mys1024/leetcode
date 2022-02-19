// 507. 完美数
// URL：https://leetcode-cn.com/problems/perfect-number/
// 难度：简单
// 关键词：数论、数学
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 99.82% 的用户

package main

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i
			j := num / i
			if i != j {
				sum += j
			}
		}
	}
	return sum == num
}
