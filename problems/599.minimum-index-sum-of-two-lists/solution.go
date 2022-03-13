// 599. 两个列表的最小索引总和
// URL：https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/
// 难度：简单
// 关键词：数组、哈希表
// 执行用时：20 ms, 在所有 Go 提交中击败了 94.25% 的用户
// 内存消耗：7 MB, 在所有 Go 提交中击败了 82.76% 的用户

package main

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i, restaurant := range list1 {
		m[restaurant] = i
	}
	ans := make([]string, 0, 64)
	minIdxSum := 999999999
	for i, restaurant := range list2 {
		idxSum, ok := m[restaurant]
		if !ok {
			continue
		}
		idxSum += i
		if idxSum > minIdxSum {
			continue
		}
		if idxSum < minIdxSum {
			minIdxSum = idxSum
			ans = ans[:0]
		}
		ans = append(ans, restaurant)
	}
	return ans
}
