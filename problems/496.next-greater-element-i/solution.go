// 496. 下一个更大元素 I
// URL：https://leetcode-cn.com/problems/next-greater-element-i/
// 难度：简单
// 关键词：栈、单调栈
// 执行用时：4 ms, 在所有 Go 提交中击败了 74.71% 的用户
// 内存消耗：3.3 MB, 在所有 Go 提交中击败了 75.29% 的用户

package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	memo := make(map[int]int)
	stack := make([]int, n)
	stackTop := 0
	for i := n - 1; i >= 0; i-- {
		num := nums2[i]
		ok := false
		for j := stackTop - 1; j >= 0; j-- {
			if stack[j] > num {
				stackTop = j + 1
				memo[num] = stack[j]
				ok = true
				break
			}
		}
		if !ok {
			stackTop = 0
			memo[num] = -1
		}
		stack[stackTop] = num
		stackTop++
	}
	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		greater := memo[num]
		ans[i] = greater
	}
	return ans
}
