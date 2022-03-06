// 6019. 替换数组中的非互质数
// URL：https://leetcode-cn.com/problems/replace-non-coprime-numbers-in-array/
// 难度：困难
// 关键词：数学、数论、最大公约数、最小公倍数、辗转相除法、链表
// 执行用时：312 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.1 MB, 在所有 Go 提交中击败了 100.00% 的用户

// 此题也可考虑使用栈来解题

package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

type LinkedListNode struct {
	Val  int
	Prev *LinkedListNode
	Next *LinkedListNode
}

func (n *LinkedListNode) ToSlice() []int {
	s := []int{}
	for n != nil {
		s = append(s, n.Val)
		n = n.Next
	}
	return s
}

func replaceNonCoprimes(nums []int) []int {
	// 使用双链表存储 nums
	head := &LinkedListNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &LinkedListNode{Val: nums[i], Prev: node}
		node = node.Next
	}
	// 在双链表中进行非互质数的替换操作
	for node := head; node.Next != nil; {
		next := node.Next
		if gcd(node.Val, next.Val) > 1 {
			node.Val = lcm(node.Val, next.Val)
			node.Next = next.Next
			next.Next = nil
			next.Prev = nil
			next = node.Next
			if next != nil {
				next.Prev = node
			}
			if node.Prev != nil {
				node = node.Prev
			}
		} else {
			node = node.Next
		}
	}
	// 返回答案
	return head.ToSlice()
}
