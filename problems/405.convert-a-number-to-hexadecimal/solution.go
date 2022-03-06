// 405. 数字转换为十六进制数
// URL：https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/
// 难度：简单
// 关键词：数学、数论
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 81.90% 的用户

package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	} else if num < 0 {
		num = 1<<32 + num
	}
	buf := []byte{}
	for num != 0 {
		digit := num % 16
		if digit < 10 {
			buf = append(buf, '0'+byte(digit))
		} else {
			buf = append(buf, 'a'+byte(digit-10))
		}
		num /= 16
	}
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
