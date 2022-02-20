// 6012. 统计各位数字之和为偶数的整数个数
// URL：https://leetcode-cn.com/problems/count-integers-with-even-digit-sum/
// 难度：简单
// 关键词：统计
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "strconv"

func countEven(num int) int {
	cnt := 0
	for i := 1; i <= num; i++ {
		sum := 0
		for _, r := range strconv.Itoa(i) {
			sum += int(r - '0')
		}
		if sum%2 == 0 {
			cnt++
		}
	}
	return cnt
}
