// 6018. 根据描述创建二叉树
// URL：https://leetcode-cn.com/problems/create-binary-tree-from-descriptions/
// 难度：中等
// 关键词：二叉树
// 执行用时：312 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.1 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type IntSet struct {
	m map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{m: map[int]bool{}}
}

func (s *IntSet) Add(num int) {
	s.m[num] = true
}

func (s *IntSet) Remove(num int) {
	delete(s.m, num)
}

func (s *IntSet) ToSlice() []int {
	slice := []int{}
	for n := range s.m {
		slice = append(slice, n)
	}
	return slice
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	rootNodes := NewIntSet()
	for _, description := range descriptions {
		parent, child := description[0], description[1]
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{Val: parent}
			rootNodes.Add(parent)
		}
		if nodes[child] == nil {
			nodes[child] = &TreeNode{Val: child}
			rootNodes.Add(parent)
		}
	}
	for _, description := range descriptions {
		parent, child, left := description[0], description[1], description[2]
		rootNodes.Remove(child)
		if left == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
	}
	root := rootNodes.ToSlice()[0]
	return nodes[root]
}
