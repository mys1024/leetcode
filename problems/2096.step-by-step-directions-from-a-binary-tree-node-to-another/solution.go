// 2096. 从二叉树一个节点到另一个节点每一步的方向
// URL：https://leetcode-cn.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
// 难度：中等
// 关键词：二叉树、深度优先搜索、哈希表
// 执行用时：308 ms, 在所有 Go 提交中击败了 42.07% 的用户
// 内存消耗：30 MB, 在所有 Go 提交中击败了 97.24% 的用户

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func repeat(b byte, count int) string {
	bytes := make([]byte, count)
	for i := 0; i < count; i++ {
		bytes[i] = b
	}
	return string(bytes)
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	// DFS 搜索起点与终点，并记录搜索路径上每个结点的父结点
	var startNode, destNode *TreeNode
	var dfs func(root *TreeNode)
	parents := map[*TreeNode]*TreeNode{root: nil}
	dfs = func(root *TreeNode) {
		if root.Val == startValue {
			startNode = root
		} else if root.Val == destValue {
			destNode = root
		}
		if startNode != nil && destNode != nil {
			return
		}
		if root.Left != nil {
			parents[root.Left] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parents[root.Right] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	// 求出根结点到起点与终点的路径
	path := func(node *TreeNode) []byte {
		bytes := []byte{}
		for parents[node] != nil {
			parent := parents[node]
			if parent.Left == node {
				bytes = append(bytes, 'L')
			} else {
				bytes = append(bytes, 'R')
			}
			node = parent
		}
		n := len(bytes)
		for i := 0; i < n/2; i++ {
			bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
		}
		return bytes
	}
	p1, p2 := path(startNode), path(destNode)
	// 找出相同的前缀路径
	idx := 0
	for idx < min(len(p1), len(p2)) {
		if p1[idx] != p2[idx] {
			break
		}
		idx++
	}
	// 求出题目要求的最短路径
	return repeat('U', len(p1)-idx) + string(p2[idx:])
}
