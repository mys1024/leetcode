// 2049. 统计最高分的节点数目
// URL：https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/
// 难度：中等
// 关键词：二叉树、深度优先搜索
// 执行用时：124 ms, 在所有 Go 提交中击败了 93.10% 的用户
// 内存消耗：23.4 MB, 在所有 Go 提交中击败了 55.17% 的用户

package main

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		parent := parents[i]
		children[parent] = append(children[parent], i)
	}
	maxScore, maxScoreCnt := 0, 0
	var dfs func(r int) int
	dfs = func(r int) int {
		allSubtreeSize, score := 0, 1
		for _, child := range children[r] {
			subtreeSize := dfs(child)
			allSubtreeSize += subtreeSize
			score *= subtreeSize
		}
		if allSubtreeSize+1 != n {
			score *= n - 1 - allSubtreeSize
		}
		if score == maxScore {
			maxScoreCnt++
		} else if score > maxScore {
			maxScore = score
			maxScoreCnt = 1
		}
		return allSubtreeSize + 1
	}
	dfs(0)
	return maxScoreCnt
}
