// 2192. 有向无环图中一个节点的所有祖先
// URL：https://leetcode-cn.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
// 难度：中等
// 关键词：图、深度优先搜索
// 执行用时：296 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：18.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	// 获取所有结点的父结点
	parents := make([][]int, n)
	for i := range parents {
		parents[i] = make([]int, 0)
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		parents[to] = append(parents[to], from)
	}
	// 定义可获取所有祖先结点的函数
	memo := make(map[int][]int, n) // getAncestors 的缓存
	var getAncestors func(int) []int
	getAncestors = func(node int) []int {
		ancestors := memo[node]
		if ancestors != nil {
			return ancestors
		}
		mapping := map[int]bool{}
		for _, p := range parents[node] {
			mapping[p] = true
			for _, a := range getAncestors(p) {
				mapping[a] = true
			}
		}
		ancestors = make([]int, len(mapping))
		i := 0
		for n := range mapping {
			ancestors[i] = n
			i++
		}
		memo[node] = ancestors
		return ancestors
	}
	// 构建答案并返回
	ans := make([][]int, n)
	for node := range ans {
		ancestors := getAncestors(node)
		sort.Slice(ancestors, func(i, j int) bool {
			return ancestors[i] < ancestors[j]
		})
		ans[node] = ancestors
	}
	return ans
}
