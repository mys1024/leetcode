// 39. 组合总和
// URL：https://leetcode-cn.com/problems/combination-sum/
// 难度：中等
// 关键词：字符串、深度优先搜索、回溯算法
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了 99.68% 的用户

package main

func combinationSum(candidates []int, target int) [][]int {
	var (
		ans         = [][]int{}
		combination = []int{}
		search      func(idx, t int)
	)

	search = func(idx, t int) {
		for i := idx; i < len(candidates); i++ {
			candidate := candidates[i]
			if candidate == t {
				cp := make([]int, len(combination)+1)
				copy(cp, combination)
				cp[len(combination)] = candidate
				ans = append(ans, cp)
			} else if candidate < t {
				combination = append(combination, candidate)
				search(i, t-candidate)
				combination = combination[:len(combination)-1]
			}
		}
	}

	search(0, target)
	return ans
}
