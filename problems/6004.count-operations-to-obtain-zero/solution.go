// 6004. 得到 0 的操作数
// URL：https://leetcode-cn.com/problems/count-operations-to-obtain-zero/
// 难度：简单
// 关键词：无
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func countOperations(num1 int, num2 int) int {
	cnt := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 -= num2
			cnt++
		} else {
			num2 -= num1
			cnt++
		}
	}
	return cnt
}
