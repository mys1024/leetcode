// 589. N 叉树的前序遍历
// URL：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
// 难度：简单
// 关键词：树、前序遍历、深度优先搜索
// 执行用时：20 ms, 在所有 Go 提交中击败了 73.93% 的用户
// 内存消耗：7.8 MB, 在所有 Go 提交中击败了 86.32% 的用户

package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}
	var dfs func(r *Node)
	dfs = func(r *Node) {
		if r == nil {
			return
		}
		ans = append(ans, r.Val)
		for _, child := range r.Children {
			dfs(child)
		}
	}
	dfs(root)
	return ans
}
