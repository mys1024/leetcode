// 6005. 使数组变成交替数组的最少操作数
// URL：https://leetcode-cn.com/problems/minimum-operations-to-make-the-array-alternating/
// 难度：中等
// 关键词：统计
// 执行用时：248 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：14.6 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import (
	"sort"
)

func minimumOperations(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	evenCounter, oddCounter := map[int]int{}, map[int]int{}
	evenFlag := true
	for _, num := range nums {
		if evenFlag {
			evenCounter[num]++
		} else {
			oddCounter[num]++
		}
		evenFlag = !evenFlag
	}

	even := [][]int{}
	for num, count := range evenCounter {
		even = append(even, []int{num, count})
	}
	sort.Slice(even, func(i, j int) bool {
		return even[i][1] > even[j][1]
	})

	odd := [][]int{}
	for num, count := range oddCounter {
		odd = append(odd, []int{num, count})
	}
	sort.Slice(odd, func(i, j int) bool {
		return odd[i][1] > odd[j][1]
	})

	if even[0][0] != odd[0][0] {
		return n - even[0][1] - odd[0][1]
	}
	if len(even) == 1 {
		if len(odd) == 1 {
			if even[0][1] > odd[0][1] {
				return odd[0][1]
			}
			return even[0][1]
		}
		return n - even[0][1] - odd[1][1]
	}
	if even[0][1] > odd[0][1] {
		return n - even[0][1] - odd[1][1]
	}
	return n - even[1][1] - odd[0][1]
}
