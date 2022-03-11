// 6018. 根据描述创建二叉树
// URL：https://leetcode-cn.com/problems/create-binary-tree-from-descriptions/
// 难度：中等
// 关键词：二叉树
// 执行用时：296 ms, 在所有 Go 提交中击败了 75.76% 的用户
// 内存消耗：8.1 MB, 在所有 Go 提交中击败了 87.37% 的用户

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	marked := map[int]bool{}

	for _, description := range descriptions {
		parent, child, isLeft := description[0], description[1], description[2]
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{Val: parent}
		}
		if nodes[child] == nil {
			nodes[child] = &TreeNode{Val: child}
		}
		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
		marked[child] = true
	}

	for nodeVal, node := range nodes {
		if !marked[nodeVal] {
			return node
		}
	}
	return nil
}
