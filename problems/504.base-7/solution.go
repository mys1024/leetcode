// 504. 七进制数
// URL：https://leetcode-cn.com/problems/base-7/
// 难度：简单
// 关键词：数学、数论
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 67.78% 的用户

package main

func convertToBase7(num int) string {
	// 特例
	if num == 0 {
		return "0"
	}
	// 准备工作
	buf := []byte{}
	negative := num < 0
	if negative {
		num = -num
		buf = append(buf, '-')
	}
	// 从低位开始提取数位
	for num != 0 {
		buf = append(buf, '0'+byte(num%7))
		num /= 7
	}
	// 翻转数字部分
	i, j := 0, len(buf)-1
	if negative {
		i = 1
	}
	for i < j {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}
