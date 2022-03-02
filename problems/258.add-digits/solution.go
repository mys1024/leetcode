// 258. 各位相加
// URL：https://leetcode-cn.com/problems/add-digits/
// 难度：简单
// 关键词：数学、数根
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

// 数学公式解法
// 求数根公式：(num-1)%9 + 1
// 参考：https://leetcode-cn.com/problems/add-digits/solution/ge-wei-xiang-jia-by-leetcode-solution-u4kj/
func addDigits(num int) int {
	return (num-1)%9 + 1
}

// // 模拟解法
// func addDigits(num int) int {
// 	for num >= 10 {
// 		sum := 0
// 		for num != 0 {
// 			tmp := num / 10
// 			sum += num - (tmp * 10)
// 			num = tmp
// 		}
// 		num = sum
// 	}
// 	return num
// }
