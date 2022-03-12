// 2191. 将杂乱无章的数字排序
// URL：https://leetcode-cn.com/problems/sort-the-jumbled-numbers/
// 难度：中等
// 关键词：数学、数论、排序
// 执行用时：332 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.2 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import (
	"math"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	numMapping := make(map[int]int, len(nums))
	for _, num := range nums {
		if num == 0 {
			numMapping[num] = mapping[0]
			continue
		}
		rawNum, mappedNum, exp := num, 0, 0
		for num != 0 {
			digit := num - (num/10)*10
			mappedNum += mapping[digit] * int(math.Pow10(exp))
			num = num / 10
			exp++
		}
		numMapping[rawNum] = mappedNum
	}
	sort.Slice(nums, func(i, j int) bool {
		return numMapping[nums[i]] < numMapping[nums[j]]
	})
	return nums
}
