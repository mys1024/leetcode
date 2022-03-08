// 67. 二进制求和
// URL：https://leetcode-cn.com/problems/add-binary/
// 难度：简单
// 关键词：数学、字符串、位运算
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 98.10% 的用户

package main

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	n := len(b) + 1
	digits := make([]byte, n)
	for i, j := len(a)-1, n-1; i >= 0; i, j = i-1, j-1 {
		digits[j] = a[i] - '0'
	}
	for i, j := len(b)-1, n-1; i >= 0; i, j = i-1, j-1 {
		if digits[j]+b[i]-'0' > 1 {
			digits[j-1]++
		}
		digits[j] = (digits[j] + b[i]) % 2
	}
	if digits[0] == 0 {
		digits = digits[1:]
	}
	for i := range digits {
		digits[i] += '0'
	}
	return string(digits)
}
