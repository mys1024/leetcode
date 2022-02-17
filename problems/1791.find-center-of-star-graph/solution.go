// 1791. 找出星型图的中心节点
// URL：https://leetcode-cn.com/problems/find-center-of-star-graph/
// 难度：简单
// 关键词：图
// 执行用时：108 ms, 在所有 Go 提交中击败了 97.56% 的用户
// 内存消耗：15.5 MB, 在所有 Go 提交中击败了 42.68% 的用户

package main

func findCenter(edges [][]int) int {
	a, b := edges[0][0], edges[0][1]
	c, d := edges[1][0], edges[1][1]
	if a == c || a == d {
		return a
	}
	return b
}
