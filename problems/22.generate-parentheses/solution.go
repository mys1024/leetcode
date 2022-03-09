// 22. 括号生成
// URL：https://leetcode-cn.com/problems/generate-parentheses/
// 难度：中等
// 关键词：字符串、深度优先搜索、回溯算法
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了 97.32% 的用户

package main

func generateParenthesis(n int) []string {
	ans := []string{}
	buf := make([]byte, 2*n)
	var search func(left, right, idx int)
	search = func(left, right, idx int) {
		if left == 0 {
			for right != 0 {
				buf[idx] = ')'
				idx++
				right--
			}
			ans = append(ans, string(buf))
		} else if left == right {
			buf[idx] = '('
			search(left-1, right, idx+1)
		} else {
			buf[idx] = '('
			search(left-1, right, idx+1)
			buf[idx] = ')'
			search(left, right-1, idx+1)
		}
	}
	search(n, n, 0)
	return ans
}
