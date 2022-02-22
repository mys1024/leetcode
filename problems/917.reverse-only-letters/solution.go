// 917. 仅仅反转字母
// URL：https://leetcode-cn.com/problems/reverse-only-letters/
// 难度：简单
// 关键词：字符串、双指针
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func isLetter(r byte) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func reverseOnlyLetters(s string) string {
	bytes := []byte(s)
	for i, j := len(s)-1, 0; i >= 0; i-- {
		b := s[i]
		if !isLetter(b) {
			continue
		}
		for !isLetter(bytes[j]) {
			j++
		}
		bytes[j] = b
		j++
	}
	return string(bytes)
}
