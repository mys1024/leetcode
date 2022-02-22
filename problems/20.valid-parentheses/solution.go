// 20. 有效的括号
// URL：https://leetcode-cn.com/problems/valid-parentheses/
// 难度：简单
// 关键词：栈
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 81.22% 的用户

package main

func isValid(s string) bool {
	var (
		brackets = map[rune]rune{
			')': '(',
			'}': '{',
			']': '[',
		}
		stack = []rune{}
		top   = 0
	)
	for _, r := range s {
		if top > 0 && brackets[r] == stack[top-1] {
			top--
		} else {
			stack = append(stack[:top], r)
			top++
		}
	}
	return top == 0
}
