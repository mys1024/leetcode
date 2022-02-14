// 1248. 统计「优美子数组」
// URL：https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
// 难度：中等
// 关键词：数组、子数组、统计
// 执行用时：92 ms, 在所有 Go 提交中击败了 97.48% 的用户
// 内存消耗：7.5 MB, 在所有 Go 提交中击败了 91.82% 的用户

// 本题用动态规划时间复杂度为 O(n^2) ，有的测试用例会超时

package main

func numberOfSubarrays(nums []int, k int) int {
	counter := make([]int, len(nums)+1)
	counter[0] = 1
	pre, ans := 0, 0
	for _, num := range nums {
		pre += num % 2
		if pre-k >= 0 {
			ans += counter[pre-k]
		}
		counter[pre]++
	}
	return ans
}
