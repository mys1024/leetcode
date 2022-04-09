// 429. N 叉树的层序遍历
// URL：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// 难度：中等
// 关键词：二叉树、广度优先搜索
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了 95.98% 的用户

package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(rootNode *Node) [][]int {
	if rootNode == nil {
		return nil
	}
	queue := []*Node{rootNode}
	ans := [][]int{}
	for len(queue) != 0 {
		size := len(queue)
		levelNodes := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes[i] = node.Val
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		ans = append(ans, levelNodes)
	}
	return ans
}
