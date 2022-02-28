// 101. 对称二叉树
// URL：https://leetcode-cn.com/problems/symmetric-tree/
// 难度：简单
// 关键词：二叉树、深度优先遍历
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3.3 MB, 在所有 Go 提交中击败了 5.58% 的用户

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nlr(root *TreeNode) []int {
	var (
		dfs      func(r *TreeNode)
		sequence = []int{}
	)
	dfs = func(r *TreeNode) {
		sequence = append(sequence, r.Val)
		if r.Left != nil {
			dfs(r.Left)
		} else {
			sequence = append(sequence, math.MaxInt)
		}
		if r.Right != nil {
			dfs(r.Right)
		} else {
			sequence = append(sequence, math.MaxInt)
		}
	}
	dfs(root)
	return sequence
}

func nrl(root *TreeNode) []int {
	var (
		dfs      func(r *TreeNode)
		sequence = []int{}
	)
	dfs = func(r *TreeNode) {
		sequence = append(sequence, r.Val)
		if r.Right != nil {
			dfs(r.Right)
		} else {
			sequence = append(sequence, math.MaxInt)
		}
		if r.Left != nil {
			dfs(r.Left)
		} else {
			sequence = append(sequence, math.MaxInt)
		}
	}
	dfs(root)
	return sequence
}

func isSymmetric(root *TreeNode) bool {
	sequence1 := nlr(root)
	sequence2 := nrl(root)
	if len(sequence1) != len(sequence2) {
		return false
	}
	for i := range sequence1 {
		if sequence1[i] != sequence2[i] {
			return false
		}
	}
	return true
}
