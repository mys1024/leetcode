// 1952. 三除数
// URL：https://leetcode-cn.com/problems/three-divisors/
// 难度：简单
// 关键词：数学
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 94.29% 的用户

package main

func isThree(n int) bool {
	if n <= 3 {
		return false
	}
	flag := false
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			if flag {
				return false
			} else {
				flag = true
			}
		}
	}
	return flag
}
