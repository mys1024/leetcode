// 974. 和可被 K 整除的子数组
// URL：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
// 难度：中等
// 关键词：前缀和、哈希表、同余定理
// 执行用时：44 ms, 在所有 Go 提交中击败了 73.81% 的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了 83.33% 的用户

// 解法参考：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/solution/he-ke-bei-k-zheng-chu-de-zi-shu-zu-by-leetcode-sol/

package main

import "fmt"

func main() {
	nums := []int{4, 5, 0, -2, -3, 1}
	k := 5
	fmt.Println(subarraysDivByK(nums, k))
}

func subarraysDivByK(nums []int, k int) int {
	record := map[int]int{0: 1}
	count, presum := 0, 0
	for _, num := range nums {
		presum += num
		remainder := presum % k
		if remainder < 0 {
			remainder = remainder + k
		}
		count += record[remainder]
		record[remainder]++
	}
	return count
}
