// 116. 填充每个节点的下一个右侧节点指针
// URL：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
// 难度：中等
// 关键词：二叉树、层序遍历
// 执行用时：4 ms, 在所有 Go 提交中击败了 92.69% 的用户
// 内存消耗：6.5 MB, 在所有 Go 提交中击败了 39.18% 的用户

package main

import (
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var (
		queue = []*Node{root}
		head  = 0
		prev  *Node
		level = 1
		i     = 1
	)
	for head != len(queue) {
		node := queue[head]
		if prev != nil {
			prev.Next = node
		}
		if i+1 == int(math.Pow(2, float64(level))) {
			prev = nil
			level++
		} else {
			prev = node
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		head++
		i++
	}
	return root
}
