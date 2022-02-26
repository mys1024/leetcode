// 606. 根据二叉树创建字符串
// URL：https://leetcode-cn.com/problems/construct-string-from-binary-tree/
// 难度：简单
// 关键词：二叉树、字符串
// 执行用时：4 ms, 在所有 Go 提交中击败了 99.10% 的用户
// 内存消耗：7.2 MB, 在所有 Go 提交中击败了 77.48% 的用户

package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) string {
	str := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return str
	}
	if root.Left != nil {
		left := dfs(root.Left)
		str += "(" + left + ")"
	} else {
		str += "()"
	}
	if root.Right != nil {
		right := dfs(root.Right)
		if right != "" {
			str += "(" + right + ")"
		}
	}
	return str
}

func tree2str(root *TreeNode) string {
	return dfs(root)
}
