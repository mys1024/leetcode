// 6024. 数组中紧跟 key 之后出现最频繁的数字
// URL：https://leetcode-cn.com/problems/most-frequent-number-following-key-in-an-array/
// 难度：简单
// 关键词：数组
// 执行用时：4 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：4.2 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func mostFrequent(nums []int, key int) int {
	mostFrequentNum, frequency := 0, 0
	cnt := map[int]int{}
	flag := false
	for _, num := range nums {
		if flag {
			cnt[num]++
			if cnt[num] > frequency {
				mostFrequentNum, frequency = num, cnt[num]
			}
		}
		flag = num == key
	}
	return mostFrequentNum
}
