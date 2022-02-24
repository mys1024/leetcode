// 537. 复数乘法
// URL：https://leetcode-cn.com/problems/complex-number-multiplication/submissions/
// 难度：中等
// 关键词：数学、字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 94.74% 的用户

package main

import "fmt"

func complexNumberMultiply(num1 string, num2 string) string {
	var a, b, c, d int
	fmt.Sscanf(num1, "%d+%di", &a, &b)
	fmt.Sscanf(num2, "%d+%di", &c, &d)
	return fmt.Sprintf("%d+%di", a*c-b*d, a*d+b*c)
}
