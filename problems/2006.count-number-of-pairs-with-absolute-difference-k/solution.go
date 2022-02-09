// 2006. 差的绝对值为 K 的数对数目
// URL：https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/
// 难度：简单
// 关键词：哈希表
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：4 MB, 在所有 Go 提交中击败了 42.00% 的用户

package main

func countKDifference(nums []int, k int) int {
	counter, memo := 0, make(map[int]int)
	for _, num := range nums {
		if k == 0 {
			counter += memo[num]
		} else {
			counter += memo[num-k] + memo[num+k]
		}
		memo[num]++
	}
	return counter
}
