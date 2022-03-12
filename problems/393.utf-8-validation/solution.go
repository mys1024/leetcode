// 393. UTF-8 编码验证
// URL：https://leetcode-cn.com/problems/utf-8-validation/
// 难度：中等
// 关键词：位运算
// 执行用时：8 ms, 在所有 Go 提交中击败了 95.45% 的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func byteType(b int) int {
	i := 7
	for i >= 0 && (b&(1<<i)) > 0 {
		i--
	}
	return 7 - i
}

func validUtf8(data []int) bool {
	n := len(data)
	for i := 0; i < n; i++ {
		bt := byteType(data[i])
		if bt == 0 {
			continue
		}
		end := bt + i
		if bt == 1 || bt > 4 || end > n {
			return false
		}
		for i++; i < end; i++ {
			if byteType(data[i]) != 1 {
				return false
			}
		}
		i--
	}
	return true
}
