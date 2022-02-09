// 703. 数据流中的第 K 大元素
// URL：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
// 难度：简单
// 关键词：堆、最小堆、优先队列、二叉树、数据流
// 执行用时：24 ms, 在所有 Go 提交中击败了 98.52% 的用户
// 内存消耗：7.8 MB, 在所有 Go 提交中击败了 91.63% 的用户

package main

// -------------------------- MinHeap Start -------------------------- //

type MinHeap struct {
	arr      []int
	size     int
	capacity int
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{make([]int, capacity+1), 0, capacity}
}

func (h *MinHeap) swap(idx1, idx2 int) {
	h.arr[idx1], h.arr[idx2] = h.arr[idx2], h.arr[idx1]
}

func (h *MinHeap) Pop() (top int) {
	if h.size == 0 {
		panic("has no element to pop")
	}
	top = h.arr[0]
	h.swap(0, h.size-1)
	h.size--
	index := 0
	for {
		left := (index+1)*2 - 1
		if left >= h.size {
			break
		}
		right := left + 1
		shouldSwap := index
		if right >= h.size {
			if h.arr[left] < h.arr[index] {
				shouldSwap = left
			}
		} else {
			minChild := left
			if h.arr[right] < h.arr[left] {
				minChild = right
			}
			if h.arr[minChild] < h.arr[index] {
				shouldSwap = minChild
			}
		}
		if shouldSwap == index {
			break
		}
		h.swap(index, shouldSwap)
		index = shouldSwap
	}
	return
}

func (h *MinHeap) Push(val int) {
	h.arr[h.size] = val
	index := h.size
	for index != 0 {
		parent := ((index + 1) / 2) - 1
		if h.arr[index] < h.arr[parent] {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
	h.size++
	if h.size > h.capacity {
		h.Pop()
	}
}

func (h *MinHeap) Top() int {
	return h.arr[0]
}

// -------------------------- MinHeap End -------------------------- //

// -------------------------- KthLargest Start -------------------------- //

type KthLargest struct {
	heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	heap := NewMinHeap(k)
	for _, num := range nums {
		heap.Push(num)
	}
	return KthLargest{heap}
}

func (kl *KthLargest) Add(val int) int {
	kl.heap.Push(val)
	return kl.heap.Top()
}

// -------------------------- KthLargest End -------------------------- //
