// 590. N 叉树的后序遍历
// URL：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
// 难度：简单
// 关键词：树、后序遍历、深度优先搜索
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了 71.51% 的用户

package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := []int{}
	var lrn func(r *Node)
	lrn = func(r *Node) {
		if r == nil {
			return
		}
		for _, child := range r.Children {
			lrn(child)
		}
		ans = append(ans, r.Val)
	}
	lrn(root)
	return ans
}
