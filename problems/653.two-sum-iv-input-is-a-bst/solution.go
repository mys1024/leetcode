// 653. 两数之和 IV - 输入 BST
// URL：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
// 难度：简单
// 关键词：二叉树、深度优先遍历
// 执行用时：20 ms, 在所有 Go 提交中击败了 85.86% 的用户
// 内存消耗：7.6 MB, 在所有 Go 提交中击败了 18.97% 的用户

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	memo := map[int]bool{}

	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		if memo[k-root.Val] {
			return true
		} else {
			memo[root.Val] = true
		}

		return dfs(root.Left) || dfs(root.Right)
	}

	return dfs(root)
}
