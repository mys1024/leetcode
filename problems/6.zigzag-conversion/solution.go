// 6. Z 字形变换
// URL：https://leetcode-cn.com/problems/zigzag-conversion/
// 难度：中等
// 关键词：字符串
// 执行用时：4 ms, 在所有 Go 提交中击败了 92.47% 的用户
// 内存消耗：3.6 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	sb := strings.Builder{}
	sb.Grow(n)
	for r := 0; r < numRows; r++ {
		offset1, offset2 := 2*(numRows-1-r), 2*(numRows-1)
		for cur := r; cur < n; cur += offset2 {
			sb.WriteByte(s[cur])
			if offset1 != 0 && offset1 != offset2 {
				if cur+offset1 >= n {
					break
				}
				sb.WriteByte(s[cur+offset1])
			}
		}
	}
	return sb.String()
}
