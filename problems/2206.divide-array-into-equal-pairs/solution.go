// 2206. 将数组划分成相等数对
// URL：https://leetcode.cn/problems/divide-array-into-equal-pairs/
// 难度：简单
// 关键词：哈希表
// 执行用时：4 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了 72.31% 的用户

package main

func divideArray(nums []int) bool {
	m := map[int]bool{}
	cnt := 0
	for _, num := range nums {
		flag := m[num]
		if flag {
			cnt++
		}
		m[num] = !flag
	}
	return cnt == len(nums)/2
}

