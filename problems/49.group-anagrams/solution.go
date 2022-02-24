// 49. 字母异位词分组
// URL：https://leetcode-cn.com/problems/group-anagrams/
// 难度：中等
// 关键词：哈希表、数组、字符串、排序
// 执行用时：20 ms, 在所有 Go 提交中击败了 86.26% 的用户
// 内存消耗：7.5 MB, 在所有 Go 提交中击败了 90.37% 的用户

package main

import "sort"

func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	var (
		ans    = make([][]string, 0, len(strs)/2)
		groups = make(map[string][]string, len(strs)/2)
	)
	for _, str := range strs {
		sorted := sortString(str)
		_, ok := groups[sorted]
		if ok {
			groups[sorted] = append(groups[sorted], str)
		} else {
			groups[sorted] = []string{str}
		}
	}
	for _, group := range groups {
		ans = append(ans, group)
	}
	return ans
}
