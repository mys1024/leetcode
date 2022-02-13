// 1108. IP 地址无效化
// URL：https://leetcode-cn.com/problems/defanging-an-ip-address/
// 难度：简单
// 关键词：字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 93.91% 的用户

package main

import "strings"

func defangIPaddr(address string) string {
	sb := strings.Builder{}
	for _, r := range address {
		if r != '.' {
			sb.WriteRune(r)
		} else {
			sb.WriteString("[.]")
		}
	}
	return sb.String()
}

// 使用正则表达式的解法
// func defangIPaddr(address string) string {
// 	reg := regexp.MustCompile(`\.`)
// 	reg.ReplaceAll([]byte(address), []byte("[.]"))
// 	return string(reg.ReplaceAll([]byte(address), []byte("[.]")))
// }
