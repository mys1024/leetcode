// 1203. 项目管理
// URL：https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/
// 难度：困难
// 关键词：图、拓扑排序、广度优先搜索
// 执行用时：48 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：10.9 MB, 在所有 Go 提交中击败了 90.00% 的用户

package main

import "sort"

// 拓扑排序
func topologicalSort(beforeItems [][]int) ([]int, bool) {
	n := len(beforeItems)
	if n == 0 {
		return []int{}, true
	}

	inDegrees := make([]int, n)
	nextItems := make([][]int, n)
	for item, beforeArr := range beforeItems {
		inDegrees[item] = len(beforeArr)
		for _, before := range beforeArr {
			if nextItems[before] == nil {
				nextItems[before] = []int{item}
			} else {
				nextItems[before] = append(nextItems[before], item)
			}
		}
	}

	queue, head, rear := make([]int, n), 0, 0
	for item, inDegree := range inDegrees {
		if inDegree == 0 {
			queue[rear] = item
			rear++
		}
	}
	if rear == 0 {
		return nil, false
	}

	for head != rear {
		item := queue[head]
		head++
		if nextItems[item] == nil {
			continue
		}
		for _, next := range nextItems[item] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue[rear] = next
				rear++
			}
		}
	}
	if rear != n {
		return nil, false
	}

	return queue, true
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// 将项目拓扑排序
	sortedItems, ok := topologicalSort(beforeItems)
	if !ok {
		return []int{}
	}

	// 将不属于现有小组的项目分配到一个独立的小组
	groupId := m
	for i, g := range group {
		if g == -1 {
			group[i] = groupId
			groupId++
			m++
		}
	}

	// 生成小组的项目表
	groupItems := make([][]int, m)
	for i, g := range group {
		if groupItems[g] == nil {
			groupItems[g] = []int{i}
		} else {
			groupItems[g] = append(groupItems[g], i)
		}
	}

	// 构建 beforeGroups
	beforeGroups := make([][]int, m)
	for i := range beforeGroups {
		beforeGroups[i] = []int{}
	}
	for item, beforeArr := range beforeItems {
		itemGroup := group[item]
		for _, before := range beforeArr {
			beforeGroup := group[before]
			if itemGroup != beforeGroup {
				beforeGroups[itemGroup] = append(beforeGroups[itemGroup], beforeGroup)
			}
		}
	}

	// 将小组拓扑排序
	sortedGroups, ok := topologicalSort(beforeGroups)
	if !ok {
		return []int{}
	}

	// 构建答案
	ans := []int{}
	itemIdx := make([]int, n)
	for idx, item := range sortedItems {
		itemIdx[item] = idx
	}
	for _, g := range sortedGroups {
		items := groupItems[g]
		for i, item := range items {
			items[i] = itemIdx[item]
		}
		sort.Ints(items)
		for i, idx := range items {
			items[i] = sortedItems[idx]
		}
		ans = append(ans, items...)
	}
	return ans
}
