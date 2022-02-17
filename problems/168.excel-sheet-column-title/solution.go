// 168. Excel表列名称
// URL：https://leetcode-cn.com/problems/excel-sheet-column-title/
// 难度：简单
// 关键词：位运算
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 95.38% 的用户

package main

func convertToTitle(columnNumber int) string {
	bytes := []byte{}
	for columnNumber != 0 {
		columnNumber--
		bytes = append(bytes, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
