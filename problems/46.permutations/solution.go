// 46. 全排列
// URL：https://leetcode-cn.com/problems/permutations/
// 难度：中等
// 关键词：回溯
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了 93.72% 的用户

package main

func permute(nums []int) [][]int {
	var (
		n           = len(nums)
		used        = make([]bool, n)
		permutation = make([]int, n)
		ans         = [][]int{}
		search      func(idx int)
	)
	search = func(idx int) {
		for i, num := range nums {
			if used[i] {
				continue
			}
			permutation[idx] = num
			if idx < n-1 {
				used[i] = true
				search(idx + 1)
				used[i] = false
			} else {
				cp := make([]int, n)
				copy(cp, permutation)
				ans = append(ans, cp)
			}
		}
	}
	search(0)
	return ans
}
