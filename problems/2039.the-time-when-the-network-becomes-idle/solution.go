// 2039. 网络空闲的时刻
// URL：https://leetcode-cn.com/problems/the-time-when-the-network-becomes-idle/
// 难度：中等
// 关键词：图、广度优先搜索
// 执行用时：356 ms, 在所有 Go 提交中击败了 33.33% 的用户
// 内存消耗：29.2 MB, 在所有 Go 提交中击败了 33.33% 的用户

package main

func networkBecomesIdle(edges [][]int, patience []int) int {
	// 构建邻接表
	adjacencies := map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjacencies[a] = append(adjacencies[a], b)
		adjacencies[b] = append(adjacencies[b], a)
	}

	// 广度优先搜索
	max := 0
	queue := [][2]int{{0, 0}}
	marked := make([]bool, len(patience))
	marked[0] = true
	for len(queue) > 0 {
		node, distance := queue[0][0], queue[0][1]
		queue = queue[1:]
		if node != 0 {
			time := 2*distance + ((2*distance-1)/patience[node])*patience[node]
			if time > max {
				max = time
			}
		}
		for _, next := range adjacencies[node] {
			if !marked[next] {
				queue = append(queue, [2]int{next, distance + 1})
				marked[next] = true
			}
		}
	}

	return max + 1
}
