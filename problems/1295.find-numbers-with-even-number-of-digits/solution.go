// 1295. 统计位数为偶数的数字
// URL：https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
// 难度：简单
// 关键词：数学
// 执行用时：4 ms, 在所有 Go 提交中击败了 93.10% 的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了 96.55% 的用户

package main

func isEven(num int) bool {
	if num == 0 {
		return false
	}
	cnt := 0
	for num != 0 {
		num /= 10
		cnt++
	}
	return cnt%2 == 0
}

func findNumbers(nums []int) int {
	cnt := 0
	for _, num := range nums {
		if isEven(num) {
			cnt++
		}
	}
	return cnt
}
