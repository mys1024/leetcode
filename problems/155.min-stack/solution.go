// 155. 最小栈
// URL：https://leetcode-cn.com/problems/min-stack/
// 难度：简单
// 关键词：栈
// 执行用时：16 ms, 在所有 Go 提交中击败了 79.53% 的用户
// 内存消耗：8 MB, 在所有 Go 提交中击败了 96.30% 的用户

package main

import "math"

type MinStack struct {
	arr []int
	top int
	min int
}

func Constructor() MinStack {
	return MinStack{
		arr: []int{},
		top: 0,
		min: math.MaxInt,
	}
}

func (s *MinStack) Push(val int) {
	if val < s.min {
		s.min = val
	}
	s.arr = append(s.arr[:s.top], val)
	s.top++
}

func (s *MinStack) Pop() {
	if s.arr[s.top-1] == s.min {
		s.min = math.MaxInt
		for i := 0; i < s.top-1; i++ {
			if s.arr[i] < s.min {
				s.min = s.arr[i]
			}
		}
	}
	s.top--
}

func (s *MinStack) Top() int {
	return s.arr[s.top-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}
