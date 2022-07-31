// 2208. 将数组和减半的最少操作次数
// URL：https://leetcode.cn/problems/minimum-operations-to-halve-array-sum/
// 难度：中等
// 关键词：数组
// 执行用时：220 ms, 在所有 Go 提交中击败了 29.41% 的用户
// 内存消耗：9.5 MB, 在所有 Go 提交中击败了 41.18% 的用户

package main

import (
	"container/heap"
)

type MaxHeap []float64

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
	h := &MaxHeap{}
	sum := float64(0)
	for _, num := range nums {
		heap.Push(h, float64(num))
		sum += float64(num)
	}
	mid := sum / 2

	heap.Init(h)
	cnt := 0
	for sum > mid {
		num := heap.Pop(h).(float64)
		num /= 2
		sum -= num
		heap.Push(h, num)
		cnt++
	}
	return cnt
}
